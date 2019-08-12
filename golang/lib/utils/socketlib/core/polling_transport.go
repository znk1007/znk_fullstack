package core

import (
	"net/http"
	"time"
)

//PollingTransport 轮询传输
type PollingTransport struct {
	Client *http.Client
}

// Default 默认
var Default = &PollingTransport{
	Client: &http.Client{
		Timeout: time.Minute,
	},
}
