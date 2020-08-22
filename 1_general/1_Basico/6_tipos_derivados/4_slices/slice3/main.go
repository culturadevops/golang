package main

/*
Ejemplo ciclos de for
En este ejemplo intento mostrar
como trabajar con slice
*/
import (
	"fmt"
	"strings"
)

func main() {
	// For slice[ i : j : k ] the
	// Length:   j - i
	// Capacity: k - i

	people := []string{
		"John",
		"Steve",
		"Jane",
	}

	p := people[1:2]
	//p := people[1:2:2] // [low:high:max] => cap = max-low
	fmt.Println("capacity:", cap(p))
	p = append(p, "Mike")

	inspect("people: ", people)
	inspect("sub people: ", p)
	fmt.Println(people[2])

	// pure clone
	peopleClone := append(people[:0:0], people...) // [0:0:0]
	inspect("people: ", people)
	inspect("peopleClone: ", peopleClone)
}

func inspect(label string, people []string) {
	f := strings.Repeat("%p | ", len(people))
	f = label + f + "\n"
	var args []interface{}
	for i := range people {
		args = append(args, &people[i])
	}
	fmt.Printf(f, args...)
}
