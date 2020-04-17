package devicemodel

import (
	usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"
)

//gormDeviceExists 设备是否存在
func gormDeviceExists(userID string) (exists bool) {
	var cnt int
	usergorm.DB().Model(&Device{
		UserID: userID,
	}).Count(&cnt)
	exists = cnt > 0
	return
}

//gormCreateDevice 创建设备信息数据
func gormCreateDevice(device Device) (exs bool, err error) {
	exs = usergorm.DB().NewRecord(device)
	if !exs {
		err = usergorm.DB().Create(device).Error
		exs = usergorm.DB().NewRecord(device)
	}
	return
}

//gormUpdateTrust 更新设备信任状态
func gormUpdateTrust(userID string, deviceID string, state DeviceState) (err error) {
	err = usergorm.DB().Model(
		&Device{
			UserID:   userID,
			DeviceID: deviceID,
		},
	).Update("state", state).Error
	return
}

//gormUpdateOnline 更新设备在线状态
func gormUpdateOnline(userID string, deviceID string, online int) (err error) {
	err = usergorm.DB().Model(
		&Device{
			UserID:   userID,
			DeviceID: deviceID,
		},
	).Update("online", online).Error
	return
}

//gormAllDevice 获取所有设备
func gormAllDevice(userID string) (devices []Device, err error) {
	var dvs []Device
	err = usergorm.DB().Model(
		&Device{
			UserID: userID,
		},
	).Find(&dvs).Error
	if err == nil {
		devices = dvs
	}
	return
}

//最近登录设备
func gormCurrentDevice(userID string) (device Device, err error) {
	var dvc Device
	err = usergorm.DB().Model(
		&Device{
			UserID: userID,
		},
	).Order("updatedAt DESC").First(&dvc).Error
	if err == nil {
		device = dvc
	}
	return
}
