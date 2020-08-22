/*
Ejemplo goto
Muestro como saltar a una etiqueta con goto

*/
package main

import "fmt"

func main() {
	var i int

CICLO:
	fmt.Println("estamos fuera del for")
	for i < 10 {
		if i == 6 {
			i = i + 3
			fmt.Println("Saltando a etiqueta CICLO con i = i + 3")
			goto CICLO2
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}
CICLO2:
	fmt.Printf("ciclo 2 Valor de i: %d\n", i)
	if i == 9 {
		fmt.Printf("Valor de i: %d\n", i)
		i = i + 3
		fmt.Println("Saltando a etiqueta CICLO con i = i + 3")
		goto CICLO
	}
	fmt.Printf("terminamos\n")
}
