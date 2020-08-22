package main

import (

	"os"
	"bufio"
	"fmt"
)

func main() {
	var (
		s string
	)
	scan:=bufio.NewScanner(os.Stdin)
	scan.Scan()
	s=scan.Text()
	fmt.Println(s)
}
