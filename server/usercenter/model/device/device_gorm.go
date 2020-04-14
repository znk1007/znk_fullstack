package devicemodel

import (
	usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"
)

//gormCreateDevice 创建设备信息数据
func gormCreateDevice(device *Device) (exs bool, err error) {
	exs = usergorm.DB().NewRecord(device)
	if !exs {
		err = usergorm.DB().Create(device).Error
		exs = usergorm.DB().NewRecord(device)
	}
	return
}

//gormUpdateTrust 更新设备信任状态
func gormUpdateTrust(userID string, deviceID string, trust int) (err error) {
	err = usergorm.DB().Model(
		&Device{
			UserID:   userID,
			DeviceID: deviceID,
		},
	).Update("trust", trust).Error
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
