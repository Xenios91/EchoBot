package Echorequest

import (
	"strings"
	"time"
)

type EchoRequest struct {
	Ip              string
	Message         string
	Token           string
	timeToTerminate *time.Time
	Delay           int
}

func (echoRequest *EchoRequest) IsPastTimeToTerminate() bool {
	var currentTime = time.Now()
	return echoRequest.timeToTerminate.Before(currentTime)
}

func (echoRequest *EchoRequest) GenerateTimeToTerminate() {
	timeToTerminate := time.Now().Add(time.Hour * time.Duration(12))
	echoRequest.timeToTerminate = &timeToTerminate
}

func (echoRequest *EchoRequest) SetPerformance(performance string) {
	performance = strings.ToLower(performance)

	switch performance {
	case "none":
		echoRequest.Delay = 0
	case "low":
		echoRequest.Delay = 2
	case "medium":
		echoRequest.Delay = 5
	case "high":
		echoRequest.Delay = 10
	case "insane":
		echoRequest.Delay = 15
	default:
		echoRequest.Delay = 0
	}
}

func (echoRequest *EchoRequest) GetPerformance() int {
	return echoRequest.Delay
}
