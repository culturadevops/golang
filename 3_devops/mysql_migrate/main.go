/*
este es la prima version de este codigo
yo ya lo mejore y separe en package
tu puedes tomar esto y mejorarlo
esta hecho para que se puede entender y mejorar
forma parte de la practica del curso de golang de youtube/culturadevops y academia virtualwebss.com
*/

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func CargarSqlFile(sqlfile string) string {

	absPath, _ := filepath.Abs(sqlfile)
	fmt.Printf(absPath)
	data, _ := ioutil.ReadFile(absPath)
	input := string(data)
	fmt.Printf(input)
	return input
}
func mysqlImport(bd bdconfig, data string) {

	buffer := bytes.Buffer{}
	buffer.Write([]byte(data))
	//mysql -u username -p database_name < file.sql

	cmd := exec.Command("mysql", "-u", bd.root,
		"-p"+bd.pass, "-h", bd.dns, bd.squema)

	cmd.Stdin = &buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}

}
func mysqlCreateSquema(bd bdconfig) {

	cmd := exec.Command("mysql", "-u", bd.root,
		"-p"+bd.pass, "-h", bd.dns, bd.squema)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}

}
func dump(bd bdconfig) {

	cmd := exec.Command("mysqldump", "--no-data", "-u", bd.root, "-p"+bd.pass, "--set-gtid-purged=OFF", "-h", bd.dns, bd.squema)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}

	CrearArchivo(bd.squema+".sql", stdout.String())

}
func CrearArchivo(rutaDestino string, data string) {
	err := ioutil.WriteFile(rutaDestino, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}

type bdconfig struct {
	dns    string
	root   string
	pass   string
	squema string
}

func main() {
	/* esto hace
	   1 dump
	   	conectarse a una bd
	   	hacer el dump
	   	guardar el archivo en una carpeta
	   	segue con otra bd hasta terminar la lista de conexion
	   2 import
	   	conectarse al destino
	   	buscar archivo
		lo sube
	   	seguir con otra import hasta terminar la lista de origen
	*/

	//array para cargar las bd que sera el origen de los dump
	/*	bdOrigen := []bdconfig{
			bdconfig{
				dns:    ".",
				root:   "root",
				pass:   "",
				squema: "",
			},
			bdconfig{
				dns:    ".",
				root:   "root",
				pass:   "",
				squema: "",
			},
			bdconfig{
				dns:    ".",
				root:   "root",
				pass:   "",
				squema: "",
			},
		}

		//aqui se hace el dump
		for key, _ := range bdOrigen {
			dump(bdOrigen[key])
		}*/
	// structura que sera el destino de la carga
	DdDestino := bdconfig{
		dns:  ".",
		root: "root",
		pass: "",
	}
	//nombre de los esquemas destinos
	NombredebdDestino := []string{
		"nombredebduno",
		"nombredebddos",
		"",
		"",
	}

	// aqui se hace la importacion recuerde que el archivo debe existir y que debes pasarle el nombre correcto

	for _, value := range NombredebdDestino {
		DdDestino.squema = value
		mysqlImport(DdDestino, CargarSqlFile(""+value+".sql"))
	}

}
