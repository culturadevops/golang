package main

import (
	"fmt"
)

func miPrimeraFuncion() {
	fmt.Println("mi funcion")
	fmt.Println("mi funcion")
	fmt.Println("mi funcion")
	fmt.Println("mi funcion")
	fmt.Println("mi funcion")
	fmt.Println("mi funcion")
}

func funcionConParametros(a string) {
	fmt.Println("mi funcion")
}
func funcionConParametrosYreturn(a int, b int) int {
	return (a + b)
}

func funcionConmultiplesParametros(arg ...string) []string {
	return arg
}
func funcionMultipleReturn(arg ...string) ([]string, string, string) {
	return arg, "err", "xxxx"
}

func main() {
	s := funcionConmultiplesParametros("1", "2", "3", "4", "5", "6")
	a, _, c := funcionMultipleReturn("1", "2", "3", "4")

	fmt.Println(s)
	fmt.Println(a)
	//fmt.Println(b)
	fmt.Println(c)

	fmt.Println("vamos a usar una funcion del sistema")
	//os.Exit(1)

	func() {
		fmt.Println("go routine 1 is done")
	}()

}
