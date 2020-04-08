package model

import (
	usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"
	userredis "github.com/znk_fullstack/server/usercenter/viewmodel/dao/redis"
)

//Device 设备信息
type Device struct {
	DeviceID string `gorm:"primary_key"`
	UserID   string
	Platform string
	Trust    int
	Online   int
}

//CreateDevice 创建设备信息数据
func CreateDevice(device *Device) (exs bool, err error) {
	exs = usergorm.DB().NewRecord(device)
	if !exs {
		err = usergorm.DB().Create(device).Error
		exs = usergorm.DB().NewRecord(device)
	}
	return
}

//UpdateTrust 更新设备信任状态
func UpdateTrust(deviceID string, trust int) error {
	return usergorm.DB().Model(&Device{DeviceID: deviceID}).Update("trust", trust).Error
}

//UpdateOnline 更新设备在线状态
func UpdateOnline(deviceID string, online int) error {
	return usergorm.DB().Model(&Device{DeviceID: deviceID}).Update("online", online).Error
}

//CurrentDevice 当前设备信息
func CurrentDevice(userID string) (deviceID string, trust int, online int) {
	deviceID = ""
	trust = 0
	online = 0
	k := userID + "_device"
	dvs, err := userredis.HMGet(k, "devicedID", "trust", "online")
	if err != nil || len(dvs) < 2 {
		return
	}
	deviceID = dvs[0].(string)
	trust = dvs[1].(int)
	online = dvs[2].(int)
	return
}

//SetCurrentDeivce 设置当前设备信息
func SetCurrentDeivce(userID string, deviceID string, trust int, online int) (e error) {
	k := userID + "_device"
	e = userredis.HSet(k, "devicedID", deviceID, "trust", trust, "online", online)
	return
}
