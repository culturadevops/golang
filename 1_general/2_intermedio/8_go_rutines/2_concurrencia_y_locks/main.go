package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
)

func setGrade(grades map[string]int, gradeName string, gradeValue int) {
	mu.Lock()
	defer mu.Unlock()
	grades[gradeName] = gradeValue
}

func main() {
	g := map[string]int{
		"English": 9,
		"Math":    8,
	}
	wg.Add(2)

	// http request 1
	go func() {
		setGrade(g, "Math", 5) // write
		wg.Done()
	}()

	// http request 2
	go func() {
		setGrade(g, "Math", 6) // write
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(g) // read
}
