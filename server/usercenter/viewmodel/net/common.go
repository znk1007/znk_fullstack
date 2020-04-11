package usernet

import (
	"github.com/rs/zerolog/log"
	"github.com/znk_fullstack/server/usercenter/model"
	userGenID "github.com/znk_fullstack/server/usercenter/viewmodel/generateId"
)

//makeID 生成唯一ID
func makeID() string {
	return userGenID.GenerateID()
}

/*
设备ID：deviceID，
平台：platform[web,iOS,Android]，
用户ID：userID，
应用标识：appkey
*/

//saveCurrentDevice 保存当前设备信息
func saveCurrentDevice(userID string, deviceID string, platform string) (err error) {
	dvs := &model.Device{
		DeviceID: deviceID,
		Platform: platform,
		Trust:    1,
		Online:   0,
		UserID:   userID,
	}
	_, err = model.CreateDevice(dvs)
	if err != nil {
		log.Info().Msg(err.Error())
		return
	}
	model.SetCurrentDeivce(userID, deviceID, 1, 0)
	return
}
