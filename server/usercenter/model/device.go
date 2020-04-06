package model

import gorm "github.com/jinzhu/gorm"

//Device 设备信息
type Device struct {
	gorm.Model
	UserID   string
	Platform string
	DeviceID string
	Trust    bool
}
