package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	//conection DB
	db, err := sql.Open("postgres", "user=postgres password=pass host=127.0.0.1
											port=5432 dbname=DBase sslmode=disable")
		if err!= nil {
		fmt.Println("sql.Open() error :  %v\n" , err)
		return
	}
defer db.Close()//Cieera la BD justo antes de concluir la ejecución
//Test conection
	err2 := db.Ping()
	if err2 != nil {
		fmt.Println("Error en prueba de conexión")
	} else {
		fmt.Println("Prueba de conexión exitoso ")
	}

	   tabla, err2 := db.Query("SELECT * FROM users")
	   if err2 !=nil {
	   	fmt.Println("Error en consulta a tabla")
	   }else {
		   for tabla.Next() {//Recorrido de la tabla
	   			var name, password string

	   			error3:=tabla.Scan(&name,&password)
	   			if error3!=nil { fmt.Println("error recorriendo la tabla")
	  				 	} else{fmt.Println(name + " " + password)}
	   		}//fin del recorrido por la tabla
	   }//end else



}//end func main
