package main

import (
	"github.com/rs/cors"
	"github.com/gorilla/mux" //go get -v -u github.com/gorilla/mux
	"net/http"
	"encoding/json"
	"log"
)


// POST
func PostCreatFileHandler(w http.ResponseWriter , r *http.Request){
	log.Println("POST para devolver el html")
	json.NewEncoder(w).Encode("<div><form><input type=\"search\"/><input type=\"search\" id=\"BBB\"/></form></div>")

	
		// log.Println("[POST] - Solicitando servicio file")
		// var documento googleDrive.Documento
		
		// //decodificamos el json recibio (request) a un objeto documento
		// error := json.NewDecoder(r.Body).Decode(&documento)
		// if error != nil {
		// 	log.Println("[POST] - Error decodificando json en Solicitando servicio file")
		// 	w.WriteHeader(http.StatusNotFound)
		// 	fmt.Fprintf(w,"Error en los parametros enviados")
		// }else{

		// 	if documento.Titulo == "" || documento.Contenido == "" {
		// 		log.Println("[POST] - Error decodificando json en Solicitando servicio file")
		// 		w.WriteHeader(http.StatusNotFound)
		// 		fmt.Fprintf(w,"Error en los parametros enviados")
		// 	}else{
		// 		resultado,status := googleDrive.CreateFile(documento)
		// 		if status != "OK"{
		// 			w.WriteHeader(http.StatusInternalServerError)
		// 			json.NewEncoder(w).Encode(status)
		// 		} else{
		// 			w.WriteHeader(http.StatusOK)
		// 			json.NewEncoder(w).Encode(resultado)
		// 		}
		// 	}
		// }

	
}

var router *mux.Router


func main() {
	router = mux.NewRouter()

    HandleFuncEx("/page", PostCreatFileHandler)

    handler := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST"}}).Handler(router)
    http.ListenAndServe(":8080", handler)

}

func HandleFuncEx(pattern string, handler func(http.ResponseWriter, *http.Request)) {
    log.Println("handled function", pattern)
    router.HandleFunc(pattern, handler)
}
