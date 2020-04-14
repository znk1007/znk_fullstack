package devicemodel

//Device 设备信息
type Device struct {
	DeviceID string `gorm:"primary_key"`
	UserID   string
	Platform string
	Trust    int
	Online   int
}
