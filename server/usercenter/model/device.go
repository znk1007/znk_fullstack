package model

import (
	usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"
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
