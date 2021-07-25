package Service

import (
	"time"
)

var (
	ticker = time.NewTicker(10 * time.Second)
)

func runServiceTasks() {
	runEchoServiceTask()
}

func executeTimer() {
	for range ticker.C {
		runServiceTasks()
	}
}

func RunTimedTask() {
	go executeTimer()
}
