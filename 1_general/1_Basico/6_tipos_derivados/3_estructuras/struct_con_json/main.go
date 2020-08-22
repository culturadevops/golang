package main

/*
Ejemplo struct
En este ejemplo intento mostrar struct con json
y como usar funciones

recuerda que este ejemplo no fue explicado en los videos
porque se intenta que lo analices y te preguntes porque pasan cosas
ejemplo
¿por que en algunos casos la estructura pasa por referencia y en otros no?
¿por que en algunos casos la estructura es retornada y en otros no?
¿puedes mejorar esta estructura?

la estructura esta hechoa para guardar items de diferentes formatos
donde el type guarda que item es y el items guarda la estructura nueva
*/
import (
	"encoding/json"
	"fmt"
)

type items struct {
	Type  string            `json:"type"`
	Items map[string]string `json:"item"`
}

func itemToken(token string) items {

	i := make(map[string]string)
	i["secret"] = token
	res := items{Type: "token", Items: i}
	return res
}

func itemCredencial(user string, pass string) items {
	i := make(map[string]string)
	i["user"] = user
	i["secret"] = pass
	res := items{Type: "credential", Items: i}
	return res
}
func itemAmazonCredencial(accountid string, user string, pass string) items {
	i := make(map[string]string)
	i["account"] = accountid
	i["user"] = user
	i["secret"] = pass
	res := items{Type: "credentialAmazon", Items: i}
	return res
}
func jsoncode(item items) []byte {

	bs1, _ := json.Marshal(item)
	fmt.Println(string(bs1))
	return bs1
}
func jsondecode(item []byte) items {
	var x items
	json.Unmarshal(item, &x)

	fmt.Println(x)
	return x
}
func main() {

	x := jsondecode(jsoncode(itemToken("pass")))
	print("\ntipo\n")
	print(x.Type)
	print("\nitem\n")
	fmt.Println(x.Items["secret"])
	//jsondecode(jsoncode(itemCredencial("user", "pass")))
	//jsondecode(jsoncode(itemAmazonCredencial("account", "user", "pass")))
}
