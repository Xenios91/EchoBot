package Echorequest

import "time"

type EchoRequest struct {
	Ip              *string
	Message         *string
	Token           *string
	timeToTerminate *time.Time
	delay           int
}

func (echoRequest *EchoRequest) IsPastTimeToTerminate() bool {
	var currentTime = time.Now()
	return echoRequest.timeToTerminate.Before(currentTime)
}

func (echoRequest *EchoRequest) GenerateTimeToTerminate() {
	timeToTerminate := time.Now().Add(time.Hour * time.Duration(4))
	echoRequest.timeToTerminate = &timeToTerminate
}

func (echoRequest *EchoRequest) SetPerformance(performance *string) {
	switch *performance {
	case "low":
		echoRequest.delay = 2
	case "medium":
		echoRequest.delay = 5
	case "high":
		echoRequest.delay = 10
	case "insane":
		echoRequest.delay = 15
	default:
		echoRequest.delay = 0
	}
}

func (echoRequest *EchoRequest) GetPerformance() int {
	return echoRequest.delay
}
