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
func CreateDevice(device *Device) (exs bool, msg string) {
	exs = usergorm.DB().NewRecord(device)
	if !exs {
		usergorm.DB().Create(device)
		exs = usergorm.DB().NewRecord(device)
		msg = "create deivce info success"
	} else {
		exs = true
		msg = "device info exists"
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
func SetCurrentDeivce(userID string, deviceID string, trust int, online int) {
	k := userID + "_device"
	userredis.HSet(k, "devicedID", deviceID, "trust", trust, "online", online)
}
