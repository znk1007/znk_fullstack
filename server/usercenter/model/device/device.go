package devicemodel

//Device 设备信息
type Device struct {
	DeviceID string `gorm:"primary_key"` //设备唯一标识
	UserID   string //用户ID
	Name     string //设备名称
	Platform string //平台
	Trust    int    //是否信任
	Online   int    //是否在线
}
