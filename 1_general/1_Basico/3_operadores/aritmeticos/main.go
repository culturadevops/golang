/*
Ejemplo de operadores aritmeticos
En este ejemplo intento
Mostrar 2 cosas;
como hace operaciones matematicas y a que me referia cuando
decia que golang es fuertemente tipado
*/
package main

import "fmt"

func main() {

	var a, b, c int

	a = 1
	b = 2
	// ¿Qué tipo de variable fue asignado a "d"?
	d := a + b
	fmt.Printf("tipo de dato d %T\n", d)
	//¿"c" no tiene decimales?
	c = a / b

	fmt.Printf("datos de variable d = %v\n", d)
	fmt.Printf("datos de variable a+b= %v\n", a+b)
	fmt.Printf("datos de variable c = %v\n", c)
	fmt.Println("-----------ahora con punto flotante ")

	var af, bf, cf float64

	af = 1.234
	bf = 2.6
	df := af + bf
	fmt.Printf("%T\n", df)
	//cf=a/bs
	cf = af / bf
	fmt.Printf("datos de variable df = %v\n", df)
	fmt.Printf("datos de variable af+bf = %v\n", af+bf)
	fmt.Printf("datos de variable cf = %v\n", cf)

	//¿Qué pasaria si "d" que es entero se asignamos una division de float64?
	df = af / bf

}
