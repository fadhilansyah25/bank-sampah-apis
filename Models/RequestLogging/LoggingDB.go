package RequestLogging

import "time"

type RequestLog struct {
	Time     time.Time `json:"time"`
	UserId   int       `json:"userid"`
	Host     string    `json:"host"`
	Method   string    `json:"method"`
	Url      string    `json:"string"`
	Status   int       `json:"status"`
	Message  string    `json:"message"`
	RemoteIp string    `json:"remoteIp"`
}
