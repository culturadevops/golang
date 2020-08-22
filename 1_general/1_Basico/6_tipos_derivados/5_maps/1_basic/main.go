package main

/*
Ejemplo map
En este ejemplo intento mostrar
como declara un map y revisar si esta nulo o no
*/
import "fmt"

func main() {
	var app map[string]int
	fmt.Println(app)
	if app == nil {
		fmt.Println("app variable is nil")
	}
}
