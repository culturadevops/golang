/*
Ejemplo Como declarar variables
En este ejemplo intento
como usar struct basico, y crear otros tipos de datos
*/
package main

import (
	"bytes"
	"fmt"
)

type mystruct struct {
	nombre    string
	segundato string
}

var miarrayint [6]int

var miarraystring [6]string

var miarrayconindice map[string]string

func main() {

	/* es posible necesitar instanciar variables de otros paquetes*/
	var stdOut bytes.Buffer
	/* asi se instancia una struct y se agrega valor*/
	var mivar = mystruct{nombre: "", segundato: "miotrodato"}
	//como instanciar un struct vacio Nota no funciona porque no se esta usando
	var mivarvacia = mystruct{}
	miarrayint[0] = 1
	miarrayint[1] = 2
	miarraystring[0] = "texto"
	miarraystring[0] = "1"
	fmt.Printf("%T", mivar)
	fmt.Println(mivarvacia)
	fmt.Println(mivar)
	miarrayconindice = make(map[string]string)
	miarrayconindice["miindice"] = "mitexto" /// falta instanciar
	fmt.Println(miarrayconindice)
	fmt.Println(miarrayint)
	fmt.Println(miarraystring)
}
