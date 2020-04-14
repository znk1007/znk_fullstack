package devicemodel

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
