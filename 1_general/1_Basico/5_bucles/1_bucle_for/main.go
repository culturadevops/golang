/*
Ejemplo ciclos de for
En este ejemplo intento mostrar muchas formas de usar el "for"
*/
package main

import "fmt"

func main() {
	var i int

	fmt.Println("-------------Primer ejemplo-----------------")
	for i <= 10 {
		fmt.Println("Valor de i:", i)
		i++
	}
	fmt.Println("-------------Segundo ejemplo-----------------")
	i = 0
	for ; i < 10; i++ {
		fmt.Println("Valor de i:", i)
	}

	fmt.Println("-------------Tercer ejemplo-----------------")

	for i = 0; i < 10; i++ {
		fmt.Println("Valor de i:", i)
	}

	fmt.Println("-------------Cuarto ejemplo-----------------")

	for i = 0; i < 10; {
		fmt.Printf("Valor de i: %d", i)
		if i == 6 {
			fmt.Printf(" sumaremos 3\n")
			i = i + 3
			continue //ejemplo para ver el continue
		}
		fmt.Printf("...\n")
		i++

	}
	fmt.Println("-------------Quinto ejemplo-----------------")
	for i = 0; i < 10; {
		fmt.Printf("Valor de i: %d", i)
		if i == 6 {
			fmt.Printf(" sumaremos 3\n")
			i = i + 3
			break //ejemplo para ver el break
		}
		fmt.Printf("...\n")
		i++
	}

	fmt.Println("-------------Sexto ejemplo-----------------")
	for i := preFor(); condicion(i); i = postFor(i) {
		fmt.Printf("Valor de i: %d", i)
		if i == 7 {
			fmt.Printf(" así que saldremos del ciclo...\n")
			break /// este ejemplo es para usar el break
		}
		fmt.Printf("\n")
	}

}
func preFor() int {
	fmt.Println("prefor i")
	return 0
}
func postFor(i int) int {
	fmt.Println("postFor sumemos i")
	i++
	return i
}
func condicion(i int) bool {
	fmt.Println("condicion i")

	return (i < 10)
}
