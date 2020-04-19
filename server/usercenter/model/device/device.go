package devicemodel

//DeviceState 设备状态
type DeviceState int

const (
	//None 待确认
	None DeviceState = iota
	//Trust 信任
	Trust DeviceState = 1
	//Reject 拒绝
	Reject DeviceState = 2
)

//Device 设备信息
type Device struct {
	DeviceID  string      `gorm:"primary_key"` //设备唯一标识
	UserID    string      //用户ID
	Name      string      //设备名称
	Platform  string      //平台
	UpdatedAt string      // 设备登录日期
	State     DeviceState //是否信任
	Online    int         //是否在线
}

//DelDevice 删除设备信息
func DelDevice(userID, deviceID string) (err error) {
	err = redisDelDevice(userID, deviceID)
	if err == nil {
		err = gormDelDevice(userID, deviceID)
	}
	return
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
func SetCurrentDevice(userID, deviceID, name, platform string, state DeviceState, update bool) (err error) {
	dvs := Device{
		DeviceID: deviceID,
		UserID:   userID,
		Name:     name,
		Platform: platform,
		State:    state,
		Online:   1,
	}
	err = redisSetCurrentDeivce(dvs)
	if err == nil {
		if update {
			if len(name) != 0 {
				err = gormUpdateName(userID, deviceID, name)
			}
			err = gormUpdateState(userID, deviceID, state)
		} else {
			_, err = gormCreateDevice(dvs)
		}
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
