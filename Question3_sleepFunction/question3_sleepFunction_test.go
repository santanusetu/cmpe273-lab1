package main

import "time"
import "testing"

func TestSleep(t *testing.T) {
	start := time.Now()
	Sleep(500)
	end := time.Now()

	if end.Sub(start) < 10 {
		t.Error("Sleep time finished before 5 second")
	}
}