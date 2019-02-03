package main

import (
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
	"os/exec"
	"./archivo"
	"time"
)

var router *mux.Router
const pathHTML string = "./webpage2html/test123.html"
const nameHTML string = "test123.html"

const nameLog string = "log.txt"

type UrlRecibido struct{  
	Url string `json:"url"`
}

type MensajeRecibido struct{  
	Url string `json:"url"`
	IdInput string `json:"idInput"`
	Contenido string `json:"contenido"`
}


// POST
func PostObtenerFileHTML(w http.ResponseWriter , r *http.Request){
	log.Println("POST para devolver el html")
	var urlRecibido UrlRecibido
	error := json.NewDecoder(r.Body).Decode(&urlRecibido)
	if error != nil {
		log.Println("[POST] - Error decodificando json en Solicitando servicio file")
		w.WriteHeader(http.StatusNotFound)
	}else{

		log.Println("Recibi la URL: ", urlRecibido.Url)
		//logica para ejecutar el script
		comando := "./buscarelhtml.sh " +  urlRecibido.Url +" "+nameHTML
		log.Println("Ejecutando el script:",comando)
		
		cmd := exec.Command("./buscarelhtml.sh", urlRecibido.Url, nameHTML)
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Esperando que el comando termine...")
		err = cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Comando termino!")
		log.Printf("Se creo el archivo .html")

		//leo el archivo
		var contenido string
		contenido = archivo.LeerArchivo(pathHTML)
		log.Println("Enviamos el response.")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(contenido)
	}
}

// POST
func PostRegistrarInput(w http.ResponseWriter , r *http.Request){
	log.Println("POST para guardar los inputs modificados")
	var mensajeRecibido MensajeRecibido
	error := json.NewDecoder(r.Body).Decode(&mensajeRecibido)
	if error != nil {
		log.Println("[POST] - Error decodificando json recibido para registrar los cambios de inputs")
		w.WriteHeader(http.StatusNotFound)
	}else{

		log.Println("Recibimos correctamente el cambio del input")

		//registramos el cambio en el json
		var mensaje string
		currentTime := time.Now()
		currentTime.Format("2006.01.02 15:04:05")

		log.Println(currentTime.String())


		mensaje = currentTime.String() +" | "+ mensajeRecibido.Url +" ID: "+ mensajeRecibido.IdInput +" CONTENT: "+ mensajeRecibido.Contenido + "\n"
		archivo.EscribirArchivo(nameLog,mensaje)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Registrado!")
	}
}

func main() {

	//eliminamos el archivo html
	archivo.BorrarArchivo(nameHTML)

	//eliminamos el archivo de log
	archivo.BorrarArchivo(nameLog)

	//creadmos el archivo de log
	archivo.CrearArchivo(nameLog)

	router = mux.NewRouter()
	log.Println("Servidor levantado.")
	
	HandleFuncEx("/logger", PostRegistrarInput)
    HandleFuncEx("/page", PostObtenerFileHTML)
	
	//se necesita configurar el cors por excepcion
    handler := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST"}}).Handler(router)
    http.ListenAndServe(":8080", handler)
}

func HandleFuncEx(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	log.Println("Tenemos el endpoint:", pattern)
	router.HandleFunc(pattern, handler)
}
