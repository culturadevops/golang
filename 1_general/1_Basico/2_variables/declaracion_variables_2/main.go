/*
Ejemplo Como declarar variables
En este ejemplo intento
mostrar como puedes averiguar que tipo de dato tiene una variable
usando el paquete fmt

para la clase en video se arreglaran errores al final del archivo*/

package main

import "fmt"

func main() {

	var variable_entera int
	var variable_string string
	variable_entera = 5
	variable_string = "texto a imprimir en pantalla"
	fmt.Printf("Tipo de datos: %T", variable_string)

	fmt.Printf("Tipo de datos: %T", variable_entera)
	fmt.Println(variable_entera)
	/*porque no se imprime uno debajo del otro que esta faltando?*/
	fmt.Printf(variable_string)

}
