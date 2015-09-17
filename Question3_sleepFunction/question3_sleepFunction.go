package main

import (
	"fmt"
	"time"
	)

func main() {	
	for i := 0; i < 11; i++ {
		fmt.Println( "->", i , "<-")
		Sleep(1000)
	}					
}

func Sleep(sec int) {
	<-time.After(time.Millisecond * time.Duration(sec))
}

