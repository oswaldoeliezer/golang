package main //iniciar archivo golang

import "net/http" //importar la libreria para manejar protocolos http

//InicioServer() funcion para iniciar el servidor http y llamar al archivo index.html
func InicioServer(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

//funcion para iniciar el servidor en la ruta "/" a travez del puerto 2016 -> este puede ser cualquiera
func main() {
	http.HandleFunc("/", InicioServer)
	http.ListenAndServe(":2016", nil)
}
