package archivo

import (
	"io/ioutil"
	"os"
	"log"
)

func LeerArchivo(path string) string{
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
	return str
}

//si esto explota mucho no importa
func BorrarArchivo(path string){
	var err = os.Remove(path)
	if err != nil { 
		log.Println("No se pudo borrar el archivo ",path) 
	}else{
		log.Println("Se borro el archivo ",path)
	}
	
}


func CrearArchivo(path string) {
	// detect if file exists
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var _, err = os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Se creo el archivo", path)
}


func EscribirArchivo(path string,mensaje string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// write some text line-by-line to file
	_, err = file.WriteString(mensaje + "\n")
	if err != nil {
		log.Fatal(err)
	}

	// save changes
	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}

  log.Println("Se guardo el registro.")
}