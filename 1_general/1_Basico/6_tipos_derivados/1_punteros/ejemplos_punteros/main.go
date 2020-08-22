package main

/*
Ejemplo de punteros
En este ejemplo intento mostrar
como usar puntero con funciones
es el inicio de la programacion orientada objeto que
se puede crear en golang
*/
import "fmt"

type people []string

func (p people) add(name string) {
	p = append(p, name)
}

func (p *people) addPtr(name string) {
	*p = append(*p, name)
}

func main() {
	world := &people{"John", "Steve"}
	world.add("Mike") // didn't work
	fmt.Println(world)
	world.addPtr("Mike") // worked
	fmt.Println(world)
}
