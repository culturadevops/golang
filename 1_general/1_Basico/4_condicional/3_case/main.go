package main

/*
Ejemplo switch case
En este ejemplo intento mostrar
un if diferente llamado switch case
es una estructura diferente que permite validar cosas
en el segundo ejemplo se muestra como reconocer el tipo de dato
de una variable interface
*/
import "fmt"

func main() {
	/*variable local de tipo entero*/
	var hora int = 10
	/*Se pasa hora como variable de prueba*/
	switch hora {
	/*Si hora coincide con alguna de las literales especificadas*/
	case 1, 2, 3, 4:
		fmt.Println("Aún es temprano")
	case 5, 6, 7:
		fmt.Println("Está atardeciendo")
	case 8:
		fmt.Println("Acaba de oscurecer")
	case 9, 10, 11:
		fmt.Println("Ya es tarde")
	default:
		fmt.Println("Es demasiado tarde")
	}
	/*variable interface sin tipo asignado*/
	var x interface{}

	switch x.(type) { /*Retorna el tipo de x*/
	/*Casos*/
	case nil:
		fmt.Println("Es una variable tipo nil")
	case int:
		fmt.Println("Es una variable tipo int")
	case float64:
		fmt.Println("Es una variable tipo float64")
	case int64:
		fmt.Println("Es una variable tipo int64")
	case string:
		fmt.Println("Es una variable tipo string")
	default:
		fmt.Println("No es ninguno de los tipos anteriores")

	}
}
