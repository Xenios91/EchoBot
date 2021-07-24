package Echorequest

import "time"

type EchoRequest struct {
	Ip              string
	Message         string
	Token           *string
	TimeToTerminate time.Time
}
