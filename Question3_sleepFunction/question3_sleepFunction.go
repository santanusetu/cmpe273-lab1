package main

import (
	"fmt"
	"time"
	)

func main() {	
	for i := 0; i < 10; i++ {
		fmt.Println( "->", i, "<-")
		Sleep(500)
	}					
}

func Sleep(sec int) {
	<-time.After(time.Millisecond * time.Duration(sec))
}

