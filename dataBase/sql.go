/*
Nombre: sql.go
Descripción: código que accede a una DB postgres y
devuelve todos los campos de una tabla users
Autor: @oswaldom876
*/
package main

import ( //bloque de importación de librerias
	"database/sql"
	"fmt"
	"log"

	//driver para postgres (debes tener internet la primera vez)
	_ "github.com/lib/pq"
)

const ( //variables constantes en bloque
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "DBase"
	sslmod   = "disable"
)

func main() {
	//Agrupamos en una variable de tipo string la información requerida
	psqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%d "+
		"dbname=%s sslmode=%s", user, password, host, port, dbname, sslmod)

	//Se abre la conexión a la BD
	db, error := sql.Open("postgres", psqlInfo)

	//evaluo el resultado de la llamada
	if error != nil { //si valor de error es distinta de nil hay un error
		fmt.Println("Error de acceso", error)
		log.Fatal(error)
	} else {
		fmt.Println("Acceso satisfactorio", error)
		log.Fatal(error)
	} //fin de la evaluación de la variable error

	//Test conection: Código para testear una conexión a la BD
	fmt.Print(">>TEST DE CONEXIÓN A LA BD: ")
	error2 := db.Ping()
	if error2 != nil {
		fmt.Println("Error en prueba de conexión!")
	} else {
		fmt.Println("Prueba de conexión exitoso! ")
	} //fin del test

	fmt.Println("..:: RESULTADO DE LA CONSULTA SQL ::.. ")
	tabla, error2 := db.Query("SELECT * FROM users")
	if error2 != nil {
		fmt.Println("Error en consulta a tabla")
	} else {
		for tabla.Next() { //Recorrido de la tabla
			var name, password string
			error3 := tabla.Scan(&name, &password)

			if error3 != nil {
				fmt.Println("error recorriendo la tabla")
			} else {
				fmt.Println(name + "\t\t" + password)
			}
		} //fin del recorrido por la tabla
	} //end else

} //end func main
