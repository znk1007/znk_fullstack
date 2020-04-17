package devicemodel

//Device 设备信息
type Device struct {
	DeviceID  string `gorm:"primary_key"` //设备唯一标识
	UserID    string //用户ID
	Name      string //设备名称
	Platform  string //平台
	UpdatedAt string // 设备登录日期
	Trust     int    //是否信任
	Online    int    //是否在线
}

//DeviceExists 设备是否存在
func DeviceExists(userID string) (exists bool) {
	exists = redisDeviceExists(userID)
	if !exists {
		exists = gormDeviceExists(userID)
	}
	return
}

//SetCurrentDevice 设置当前设备
func SetCurrentDevice(userID, deviceID, name, platform string, trust int) (err error) {
	dvs := Device{
		DeviceID: deviceID,
		UserID:   userID,
		Name:     name,
		Platform: platform,
		Trust:    trust,
		Online:   1,
	}
	err = redisSetCurrentDeivce(dvs)
	if err == nil {
		_, err = gormCreateDevice(dvs)
	}
	return
}

//CurrentDevice 当前设备信息
func CurrentDevice(userID string) (device Device, err error) {
	device, err = redisCurrentDevice(userID)
	if err != nil {
		device, err = gormCurrentDevice(userID)
	}
	return
}
