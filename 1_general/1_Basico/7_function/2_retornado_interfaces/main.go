package main

import "fmt"

type person struct {
}

type items struct {
}

func main() {
	// type & value MUST BE NIL
	p := adult(19)

	switch p.(type) {

	case nil:
		fmt.Println("Es una variable tipo nil")
	case int:
		fmt.Println("Es una variable tipo int")
	case float64:
		fmt.Println("Es una variable tipo float64")
	case int64:
		fmt.Println("Es una variable tipo int64")
	case string:
		fmt.Println("Es una variable tipo string")
	case person:
		fmt.Println("Es una variable tipo person")
	case items:
		fmt.Println("Es una variable tipo items")
	default:
		fmt.Println("No es ninguno de los tipos anteriores")

	}
	/*
		if p != nil {
			fmt.Println(":(, I'm still a kid")
		} else {
			fmt.Println("He-he, finally got adult")
		}
	*/
}

func adult(n int) interface{} {
	//var res *person = nil
	res := person{}

	if n < 18 {
		return nil // type: nil | value: nil
	}
	return res // type: *person | value: nil
}
