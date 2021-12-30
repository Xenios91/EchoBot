package Echorequest

import (
	"strings"
	"time"
)

type EchoRequest struct {
	IP              string
	Message         string
	Token           string
	timeToTerminate *time.Time
	Delay           int
	ContentType     string
}

func (echoRequest *EchoRequest) IsPastTimeToTerminate() bool {
	var currentTime = time.Now()
	return echoRequest.timeToTerminate.Before(currentTime)
}

func (echoRequest *EchoRequest) GenerateTimeToTerminate() {
	timeToTerminate := time.Now().Add(time.Hour * time.Duration(12))
	echoRequest.timeToTerminate = &timeToTerminate
}

func (echoRequest *EchoRequest) SetContentType(contentType string) {
	contentType = strings.ToLower(contentType)

	switch contentType {
	case "json":
		echoRequest.ContentType = "application/json"
	case "xml":
		echoRequest.ContentType = "application/xml"
	case "plaintext":
		echoRequest.ContentType = "text/plain"
	default:
		echoRequest.ContentType = "text/plain"
	}
}

func (echoRequest *EchoRequest) SetPerformance(performance string) {
	performance = strings.ToLower(performance)

	switch performance {
	case "none":
		echoRequest.Delay = 0
	case "low":
		echoRequest.Delay = 1
	case "medium":
		echoRequest.Delay = 3
	case "high":
		echoRequest.Delay = 7
	case "insane":
		echoRequest.Delay = 10
	default:
		echoRequest.Delay = 0
	}
}
