package main

import "time"
import "testing"

func TestSleep(t *testing.T) {
	startTime := time.Now()
	Sleep(1000)
	endTime := time.Now()

	if endTime.Sub(startTime) < 1000 {
		t.Error("Sleep time finished before 1 second")
	}
}