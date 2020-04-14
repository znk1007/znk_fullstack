package devicemodel

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
