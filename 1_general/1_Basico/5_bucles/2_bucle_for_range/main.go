/*
Ejemplo ciclos de for
En este ejemplo intento mostrar muchas formas de usar el "for range"
*/
package main

import "fmt"

func main() {

	arreglo := [8]int{0, 1, 4, 6, 10, 9}
	for i, j := range arreglo {
		fmt.Printf("indice i: %d tiene como valor #%d\n", i, j)
	}
	for i := range arreglo {
		fmt.Printf("Valor de i: %d\n", i)
	}
	for _, j := range arreglo {
		fmt.Printf("Valor de i: %d\n", j)
	}

}
