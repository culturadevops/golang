package main

/*
Ejemplo struct
En este ejemplo intento mostrar como trabajar
con estructuras anomimas
*/
//https://repl.it/@steevehook/anonymous-structs#main.go
import (
	"encoding/json"
	"fmt"
)

func main() {
	//ejemplo de un struct anomino
	res := struct {
		Name string `json:"name_of"`
		Age  int    `json:"age_of"`
	}{ //con su declaracion de datosss
		Name: "Steve",
		Age:  25,
	}
	fmt.Println(res)
	bs, _ := json.Marshal(res)
	fmt.Println(bs)
	fmt.Println(string(bs))
}
