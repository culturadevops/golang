package main

import (
	"crypto/sha256"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func Getchecksum(fileName string) string {

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return string(h.Sum(nil))

}

type misfiles struct {
	Fotos      map[string]string
	Duplicadas map[string]string
	//count      int
}

func (mf *misfiles) MySearchFiles(rutaFinal string) {
	//fmt.Println(mf.count)
	//fmt.Println(path)
	if _, err := os.Stat(rutaFinal); !os.IsNotExist(err) {

		files, err := ioutil.ReadDir(rutaFinal)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			//mf.count++
			path := rutaFinal + "/" + f.Name()

			if f.IsDir() {

				mf.MySearchFiles(path)

			} else {
				//for _, value := range format {
				//	if strings.Contains(f.Name(), value) {
				checksum := Getchecksum(path)
				if _, ok := mf.Fotos[checksum]; ok {
					mf.Duplicadas[path] = f.Name()
				} else {
					mf.Fotos[checksum] = path

				}
				//	}
				//	}
			}

		}
	}
}

func main() {
	t := time.Now()
	rutaFinal := "origen"
	//formats := []string{".gif", ".jpg", ".jpeg"}
	//fotos := make(map[string]string)
	//duplicadas := make(map[string]string)
	f := &misfiles{
		Fotos:      make(map[string]string),
		Duplicadas: make(map[string]string),
		//count:      0,
	}
	//var count *int
	f.MySearchFiles(rutaFinal)
	for path, value := range f.Duplicadas {
		err := os.Rename(path, "duplicados/"+string(t.Unix())+"-"+value)
		if err != nil {
			panic(err)
		}
	}
}
