package devicemodel

const (
	devicePrefix = "current_device"
)

//redisCurrentDevice redis中当前设备信息
func redisCurrentDevice(userID string) (deviceID string, trust int, online int) {
	deviceID = ""
	trust = 0
	online = 0
	key := devicePrefix + userID
	dvs, err := userredis.HMGet(key, "devicedID", "trust", "online")
	if err != nil || len(dvs) < 2 {
		return
	}
	deviceID = dvs[0].(string)
	trust = dvs[1].(int)
	online = dvs[2].(int)
	return
}

//redisSetCurrentDeivce redis设置当前设备信息
func redisSetCurrentDeivce(userID string, deviceID string, trust int, online int) (e error) {
	key := devicePrefix + userID
	e = userredis.HSet(key, "devicedID", deviceID, "trust", trust, "online", online)
	return
}
