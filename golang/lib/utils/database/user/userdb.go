package userdb

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/znk1007/fullstack/lib/utils/database/cockroachz"
	"github.com/znk1007/fullstack/lib/utils/database/redisz"
	"github.com/znk1007/fullstack/lib/utils/database/redisz/redigo/redis"
)

// User 用户管理客户端
type User struct {
	gorm.Model
	UserID   string `redis:"userId" gorm:"column:userid; primary_key:true"`
	Account  string `redis:"account" gorm:"column:account"`
	Nickname string `gorm:"column:nickname"`
	Phone    string `gorm:"column:phone"`
	Email    string `gorm:"column:email"`
	Photo    string `gorm:"column:photo"`
	Password string `json:"password" gorm:"column:password"`
	Active   bool   `json:"active" gorm:"column:active"`
	Device   string `json:"device" gorm:"column:device"`
	Online   bool   `json:"online" gorm:"column:online"`
	Admin    bool   `json:"admin" gorm:"column:admin"`
}

func init() {
	cockroachz.Client.AutoMigrate(&User{})
}

// Insert 插入用户数据
func (u User) Insert() (int32, error) {
	if cockroachz.Client.NewRecord(&u) {
		err := cockroachz.Client.Create(&u).Error
		if err != nil {
			return 0, err
		}
	}

	if !userExistsInMemery(u.Account) {
		rm := map[string]string{
			"account": u.Account,
			"userId":  u.UserID,
		}
		err := redisz.Manager.HMSet(redis.Args{}.Add(u.Account).AddFlat(rm)...)
		if err != nil {
			return 0, err
		}
	}

	return 1, nil
}

// GetActiveUser 获取活跃状态用户
func (u *User) GetActiveUser(userID string) error {
	err := cockroachz.Client.First(&u, "userId = ? AND active = ?", userID, true).Error
	return err
}

// IsUserExists 用户是否存在
func IsUserExists(account string) (bool, error) {
	if account == "" {
		return false, errors.New("account cannot be empty")
	}
	existsInMemery := userExistsInMemery(account)
	existsInDisk := userExistsInDisk(account)
	return (existsInMemery && existsInDisk), nil
}

// userExistsInMemery 是否存在于redis缓存中
func userExistsInMemery(account string) bool {
	exists, err := redisz.Manager.Exists(account)
	if err == nil {
		return exists != 0
	}
	return true
}

// userExistsInDisk 是否存在磁盘
func userExistsInDisk(account string) bool {
	userID, err := getUserIDFromDisk(account)
	if err != nil {
		return false
	}
	return userID != ""
}

// getUserIDFromDisk 从磁盘获取用户ID
func getUserIDFromDisk(account string) (string, error) {
	user := User{}
	err := cockroachz.Client.First(&user, "account = ? AND active = ?", account, true).Error
	if err != nil {
		return "", err
	}
	return user.UserID, nil
}

// GetSessionID 获取会话ID
func GetSessionID(userID string) string {
	exists, _ := redisz.Manager.Exists(userID)
	sessionID, _ := redisz.Manager.GetString(userID)
	if exists == 0 || sessionID == "" || sessionID == "default" {
		return ""
	}
	return sessionID
}

// GetUserID 获取用户ID
func GetUserID(account string) (string, error) {
	u := User{}
	err := redisz.Manager.HGetAllStructValue(account, &u)
	if err == nil {
		return u.UserID, nil
	}
	var userID string
	userID, err = getUserIDFromDisk(u.Account)
	if err != nil {
		return "", err
	}
	return userID, nil
}

// UpdateSessionIDAndOnlineState 更新sessionId和在线状态
func UpdateSessionIDAndOnlineState(userID string, sessionID string, timeoutSeconds int, online bool) error {
	redisz.Manager.SetStringWithExpire(userID, sessionID, timeoutSeconds)
	fmt.Println("update online: ", online)
	user := User{}
	err := cockroachz.Client.Model(&user).Where("active = ? AND userId = ?", true, userID).Update("online", online).Error
	return err
}

// IsUserActive 用户是否活跃
func IsUserActive(userID string) (bool, error) {
	if userID == "" {
		return false, errors.New("userID cannot be empty")
	}
	u := User{}
	err := cockroachz.Client.First(&u, "userId = ?", userID).Error
	return u.Active, err
}

// IsUserOnline 用户是否在线
func IsUserOnline(userID string) (bool, error) {
	if userID == "" {
		return false, errors.New("userID cannot be empty")
	}
	u := User{}
	err := cockroachz.Client.First(&u, "userId = ?", userID).Error
	return u.Online, err
}

// UpdateOnline 更新在线状态
func UpdateOnline(userID string, sessionID string, online bool) error {
	sID := GetSessionID(userID)
	if userID == "" || sessionID == "" {
		return errors.New("userID and sessionID cannot be empty")
	}
	if sID != sessionID {
		return errors.New("sessionID invalid, please sign in again")
	}
	active, err := IsUserActive(userID)
	if err != nil || active == false {
		return errors.New("user is inactive")
	}
	return cockroachz.Client.Model(&User{}).Where("userId = ? AND active = ?", userID, true).Update("online", online).Error
}

// UpdatePhoto 更新头像
func UpdatePhoto(userID string, sessionID string, photo string) error {
	sID := GetSessionID(userID)
	if userID == "" || sessionID == "" {
		return errors.New("userID and sessionID cannot be empty")
	}
	if sID != sessionID {
		return errors.New("sessionID invalid, please sign in again")
	}
	active, err := IsUserActive(userID)
	if err != nil || active == false {
		return errors.New("user is inactive")
	}
	err = cockroachz.Client.Model(&User{}).Where("userId = ? AND active = ?", userID, true).Update("photo", photo).Error
	return err
}

// UpdateNickname 更新昵称
func UpdateNickname(userID string, sessionID string, nickname string) error {
	sID := GetSessionID(userID)
	if userID == "" || sessionID == "" {
		return errors.New("userID and sessionID cannot be empty")
	}
	active, err := IsUserActive(userID)
	if err != nil || active == false {
		return errors.New("user is inactive")
	}

	if sID != sessionID {
		return errors.New("sessionID invalid, please sign in again")
	}

	err = cockroachz.Client.Model(&User{}).Where("userId = ? AND active = ?", userID, true).Update("nickname", nickname).Error
	return err
}

// IsAdmin 是否为管理员
func IsAdmin(userID string) (bool, error) {
	if userID == "" {
		return false, errors.New("userID cannot be empty")
	}
	u := User{}
	err := cockroachz.Client.First(&u, "userId = ?", userID).Error
	return u.Admin, err
}

// ActiveUser 开启/禁用用户
func ActiveUser(userID string, accounts []string, active bool) error {
	admin, err := IsAdmin(userID)
	if err != nil || admin == false {
		return errors.New("user is not the admin")
	}

	removeFromMemoery(accounts)

	err = cockroachz.Client.Model(&User{}).Where("account in (?)", accounts).Update("active", active).Error
	return err
}

// Remove 物理删除账号
func Remove(userID string, accounts []string) error {
	admin, err := IsAdmin(userID)
	if err != nil || admin == false {
		return errors.New("user is not the admin")
	}
	removeFromMemoery(accounts)
	err = cockroachz.Client.Where("account in (?)", accounts).Unscoped().Delete(User{}).Error
	return err
}

// 从缓存中删除
func removeFromMemoery(accounts []string) {
	for _, account := range accounts {
		u := User{}
		err := redisz.Manager.HGetAllStructValue(account, &u)
		if err != nil {
			_, err = redisz.Manager.Del(u.UserID)
		}
		_, err = redisz.Manager.Del(account)

	}

}
