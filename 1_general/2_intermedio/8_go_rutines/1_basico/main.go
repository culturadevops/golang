package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		fmt.Println("1")
	}()
	go func() {
		fmt.Println("2")
	}()
	time.Sleep(time.Second)
}
