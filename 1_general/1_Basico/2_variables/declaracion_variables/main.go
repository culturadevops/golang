/*
Ejemplo Como declarar variables
En este ejemplo intento
mostrar todas las formas de declaración que acepta golang

el archivo NO va a correr porque no fue terminado...

te invito a terminarlo usa todas las variables de alguna forma.
Este ejemplo no correra tiene errores
*/

package main

import "fmt"

//import "fmt"

/* funcion principal*/
func main() {
	var (
		i int
		b bool
		s string
	)
	/*FORMA TRADICIONA*/
	var variable_entera int
	var variable_string string

	/*FORMA ITERATIVA*/
	var entero = 5
	otroentero := 6

	/*FORMA EN GRUPO*/
	var x, y, z int
	i = 1
	b = true
	s = "un texto cualquiera"
	fmt.Println(i)
	fmt.Println(b)
	fmt.Println(s)

	variable_entera = 5
	variable_string = "otro texto"
	fmt.Println(variable_entera)
	fmt.Println(variable_string)
	/*FORMA ITERATIVA*/
	fmt.Println(entero)
	fmt.Println(otroentero)

	/*FORMA EN GRUPO*/
	x = 10
	y = 44
	z = 66
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	/*
	   para hacer funcionar este archivo debe usar todas la variables,
	   te recomiendo que imprimas en pantallas cada una de las variables declaradas anteriormente,
	   no olvides asignarle algun valor
	*/

	/* NOTA IMPORTANTE SIEMPRE DEBES USAR LAS VARIABLES QUE HAS DECLARADO*/

}
