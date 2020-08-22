package main

/*
Ejemplo de puntero
En este ejemplo intento mostrar
como apuntar a variables string
*/
import "fmt"

func main() {
	message := "Hello Cosmos"

	// Pointer to string
	var pMessage *string

	// pMessage points to addr of message
	pMessage = &message
	fmt.Println("Message = ", *pMessage)
	fmt.Println("Message Address = ", pMessage)

	// Change message using pointer de-referencing
	*pMessage = "Hello Universe"
	fmt.Println("Message = ", *pMessage)
	fmt.Println("Message Address = ", &message)
}
