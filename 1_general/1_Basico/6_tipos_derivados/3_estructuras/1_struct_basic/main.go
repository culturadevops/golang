package main

/*
Ejemplo de struct
En este ejemplo intento mostrar
el trabajo con struct
*/
import "fmt"

type items struct {
	Type  string
	Items map[string]string
}

func main() {
	var items_vacio = items{}
	items_vacio.Type = "text"
	items_vacio.Items = make(map[string]string)
	items_vacio.Items["item1"] = "item1tipo1"
	i := make(map[string]string)
	i["item1"] = "item1tipo1"
	i["item2"] = "item1tipo1"
	i["item3"] = "item1tipo1"
	i["item1"] = "item1tipo1xxxxxxxxx"

	items_con_algo := items{Type: "tipo1", Items: i}
	fmt.Println(items_vacio)
	fmt.Println(items_con_algo)
}
