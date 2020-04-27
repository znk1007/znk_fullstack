package devicemodel

import (
	"errors"

	userredis "github.com/znk_fullstack/server/usercenter/source/controller/dao/redis"
)

const (
	devicePrefix = "current_device"
)

//redisDelDevice 删除设备
func redisDelDevice(userID, deviceID string) (err error) {
	dID, _ := redisCurrentDeviceID(userID)
	if dID == deviceID {
		key := devicePrefix + userID
		err = userredis.Del(key)
	}
	return
}

//redisDeviceExists 用户是否存在关联设备
func redisDeviceExists(userID string) (exists bool) {
	key := devicePrefix + userID
	exists = userredis.Exists(key)
	return
}

//redisCurrentDevice 获取当前设备信息
func redisCurrentDevice(userID string) (device Device, err error) {
	key := devicePrefix + userID
	var vals []interface{}
	vals, err = userredis.HMGet(
		key,
		"deviceID",
		"state",
		"platform",
		"name",
		"userID",
		"updatedAt",
	)
	deviceID, _ := vals[0].(string)
	state, _ := vals[1].(DeviceState)
	platform, _ := vals[2].(string)
	name, _ := vals[3].(string)
	orgUserID, _ := vals[4].(string)
	updatedAt, _ := vals[5].(string)
	if orgUserID != userID {
		err = errors.New("user not match")
		return
	}
	device = Device{
		DeviceID:  deviceID,
		UserID:    userID,
		Name:      name,
		Platform:  platform,
		UpdatedAt: updatedAt,
		State:     state,
		Online:    1,
	}
	return
}

//redisCurrentDeviceID redis中当前设备部分信息
func redisCurrentDeviceID(userID string) (deviceID string, state int) {
	deviceID = ""
	state = 0
	key := devicePrefix + userID
	dvs, err := userredis.HMGet(key, "deviceID", "state")
	if err != nil || len(dvs) < 2 {
		return
	}
	deviceID, _ = dvs[0].(string)
	state, _ = dvs[1].(int)
	return
}

//redisSetCurrentDeivce redis设置当前设备信息
func redisSetCurrentDeivce(device Device) (e error) {
	key := devicePrefix + device.UserID
	e = userredis.HSet(
		key,
		"deviceID", device.DeviceID,
		"state", device.State,
		"online", device.Online,
		"platform", device.Platform,
		"name", device.Name,
		"userID", device.UserID,
		"updatedAt", device.UpdatedAt,
	)
	return
}

//redisUpdateDeviceState 更新信任设备
func redisUpdateDeviceState(userID, deviceID string, state DeviceState) (e error) {
	key := devicePrefix + userID
	orgDID, _ := redisCurrentDeviceID(userID)
	if orgDID != deviceID {
		return errors.New("not the same device")
	}
	e = userredis.HSet(key, "state", state)
	return
}

//redisUpdateDeviceOnline 更新设备在线状态
func redisUpdateDeviceOnline(userID, deviceID string, online int) (e error) {
	key := devicePrefix + userID
	orgDID, _ := redisCurrentDeviceID(userID)
	if orgDID != deviceID {
		return errors.New("not the same device")
	}
	e = userredis.HSet(key, "online", online)
	return
}
