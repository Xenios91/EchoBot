package Service

import (
	echoRequest "EchoBot/echo_request"
	"crypto/rand"
	"fmt"
	"sync"
)

type EchoRequestService struct {
	echoMap map[string]echoRequest.EchoRequest
}

var (
	doOnce             sync.Once
	echoRequestService *EchoRequestService
)

func runEchoServiceTask() {
	echoServiceCleanUpTask()
}

func echoServiceCleanUpTask() {
	if echoRequestService != nil {
		for key, value := range echoRequestService.echoMap {
			if value.IsPastTimeToTerminate() {
				delete(echoRequestService.echoMap, key)
			}
		}
	}
}

func generateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	token := fmt.Sprintf("%x", b)
	if _, ok := echoRequestService.echoMap[token]; ok {
		token = generateToken()
	}
	fmt.Println(len(token))
	return token
}

func GetEchoRequestService() *EchoRequestService {
	once := func() {
		echoRequestService = &EchoRequestService{}
		echoRequestService.echoMap = make(map[string]echoRequest.EchoRequest)
	}
	doOnce.Do(once)
	return echoRequestService
}

func (echoRequestService EchoRequestService) GetEchoMap() map[string]echoRequest.EchoRequest {
	return echoRequestService.echoMap
}

func (echoRequestService EchoRequestService) AddToMap(echoRequest *echoRequest.EchoRequest) string {
	echoRequest.Token = generateToken()
	echoRequest.GenerateTimeToTerminate()
	echoRequestService.echoMap[echoRequest.Token] = *echoRequest
	return echoRequest.Token
}
