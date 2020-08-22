/*
Ejemplo de condicionales basicos
En este ejemplo intento mostrar
como hacer if, else , else if
*/
package main

import "fmt"

func main() {
	var a, b int
	a = 1
	b = 2

	if a == 1 {
		fmt.Println("a=1")
	} else {
		fmt.Println("a!=1")
	}

	if b > 1 {
		fmt.Println("b>1")
	}
	calificacion := 10
	// ejemplo de else if
	if calificacion < 6 {
		fmt.Println("Reprobaste.")
	} else if calificacion >= 6 && calificacion <= 8 {

		fmt.Println("Aprobaste.")
	} else if calificacion == 9 {

		fmt.Println("Aprobaste. Te fue muy bien.")
	} else {

		fmt.Println("Felicidades. Aprobaste con calificación perfecta")
	}
	fmt.Println("Tu calificación fue de: ", calificacion)

}
