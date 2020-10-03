/*
Ejemplo if complejo
En este ejemplo intento mostrar
un if diferente que puedes conseguir en ejemplos
de la documentación oficial y en foros de internet
*/
package main

import (
	"fmt"
	"os"
)

func main() {

	var rutaAbuscar string
	rutaAbuscar = "main.g1o"

	if _, err := os.Stat(rutaAbuscar); !os.IsNotExist(err) {
		fmt.Println("existe")
	} else {
		fmt.Println("no existe")
	}
	rutaAbuscar = "main.go"
	if _, err := os.Stat(rutaAbuscar); !os.IsNotExist(err) {
		fmt.Println("existe")
	} else {
		fmt.Println("no existe")
	}
}
