package main

/*
Ejemplo punteros
En este ejemplo intento mostrar como usar de forma basica
los punteros
el signo * y el signo & c
como podemos ver la direccion de memoria de la variable original y
la direccion de memoria del puntero
*/
import "fmt"

func main() {
	var mi_var = 100
	var dir_var *int
	dir_var = &mi_var //puntero apunta a una variable

	fmt.Printf("Valor de la variable 'mi_var': %d \n", mi_var)
	fmt.Printf("Dirección almacenada en 'dir_var': %x \n", dir_var)
	fmt.Printf("Valor de la variable que apunta 'dir_var': %d \n", *dir_var)
	fmt.Printf("Dirección que ocupa el apuntador 'dir_var' en memoria: %x \n", &dir_var)

	fmt.Printf("Dirección  en memoria de mi_var: %x \n", &mi_var)
	//sumaVariableNormal(mi_var)
	//fmt.Printf("Valor de la variable 'mi_var': %d \n", mi_var)
	//como cambiamos el valor de la variable?
	sumaPorReferencia(&mi_var)
	fmt.Printf("Valor de la variable 'mi_var': %d \n", mi_var)
	//no lo hagas
	//sumaPorReferenciaunPuntero(&dir_var)
	//fmt.Printf("Valor de la variable 'mi_var': %d \n", mi_var)
}
func sumaPorReferencia(a *int) {
	*a = *a + 10
	fmt.Printf("Dirección que ocupa variable 'a': %x \n", a)
	//	fmt.Printf("Dirección almacenada en 'a': %x \n", a)
}
func sumaVariableNormal(a int) {
	a = a + 10
	fmt.Printf("Dirección que ocupa variable 'a': %x \n", &a)

}

func sumaPorReferenciaunPuntero(a **int) {
	**a = **a + 10
	//fmt.Printf("Dirección que ocupa variable 'a': %x \n", &a)
	//	fmt.Printf("Dirección almacenada en 'a': %x \n", a)
}
