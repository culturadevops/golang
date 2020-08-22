package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(struct{}{})) // 0
	done := make(chan struct{})
	go func() {
		fmt.Println("go routine 1 is done")
		done <- struct{}{}
	}()
	go func() {
		fmt.Println("go routine 2 is done")
		done <- struct{}{}
	}()

	for {
		select {
		case <-done:
			fmt.Println("done at", time.Now())
		case <-time.After(time.Second):
			fmt.Println("done waiting")
			return
		}
	}
}
