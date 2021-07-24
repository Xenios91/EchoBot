package Service

import (
	echoRequest "EchoBot/echo_request"
	"crypto/rand"
	"fmt"
	"sync"
	"time"
)

type EchoRequestService struct {
	echo_map map[string]echoRequest.EchoRequest
}

var (
	doOnce             sync.Once
	echoRequestService *EchoRequestService
)

func generateToken() *string {
	b := make([]byte, 12)
	rand.Read(b)
	token := fmt.Sprintf("%x", b)
	if _, ok := echoRequestService.echo_map[token]; ok {
		token = *generateToken()
	}
	return &token
}

func GetEchoRequestService() *EchoRequestService {
	once := func() {
		echoRequestService = &EchoRequestService{}
		echoRequestService.echo_map = make(map[string]echoRequest.EchoRequest)
	}

	doOnce.Do(once)
	return echoRequestService
}

func (echoRequestService EchoRequestService) GetEchoMap() map[string]echoRequest.EchoRequest {
	return echoRequestService.echo_map
}

func (echoRequestService EchoRequestService) AddToMap(echoRequest echoRequest.EchoRequest) *string {
	echoRequest.Token = generateToken()
	echoRequest.TimeToTerminate = time.Now().Local().Add(time.Hour * time.Duration(4))
	echoRequestService.echo_map[*echoRequest.Token] = echoRequest
	return echoRequest.Token
}
