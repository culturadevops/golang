package main

/*
Ejemplo struct
En este ejemplo intento mostrar struct con funciones

*/
import "fmt"

//https://repl.it/@steevehook/type-aliases#main.go

type human struct {
	name string
	age  int
}

func (h *human) setName(n string) {
	h.name = n
}
func (h *human) setAge(a int) {
	h.age = a
}

//asi se puede copiar una structura y sus funciones
type student = human

func info(h human) {
	fmt.Printf("Hi my name is: %s and I'm: %d\n", h.name, h.age)
}

func main() {
	h := human{name: "John", age: 25}
	s := student{}
	s.setName("Jane")
	s.setAge(23)
	info(h)
	info(s)
}
