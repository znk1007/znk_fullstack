package devicemodel

import (
	userredis "github.com/znk_fullstack/server/usercenter/viewmodel/dao/redis"
)

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
func redisSetCurrentDeivce(device Device) (e error) {
	key := devicePrefix + device.UserID
	e = userredis.HSet(
		key,
		"devicedID", device.DeviceID,
		"trust", device.Trust,
		"online", device.Online,
		"platform", device.Platform,
		"name", device.Name,
		"userId", device.UserID,
	)
	return
}

//redisUpdateDeviceTrust 更新信任设备
func redisUpdateDeviceTrust(userID, deviceID string, trust int) (e error) {
	key := devicePrefix + userID
	e = userredis.HSet(key, "trust", trust)
	return
}

//redisUpdateDeviceOnline 更新设备在线状态
func redisUpdateDeviceOnline(userID, deviceID string, online int) (e error) {
	key := devicePrefix + userID
	e = userredis.HSet(key, "online", online)
	return
}
