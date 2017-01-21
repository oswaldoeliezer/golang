package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq/oid"
)

func main() {
	bd, err := sql.Open("pq", "user='postgres' dbname='DBase' password='pass' host='127.0.0.1' port='5432'")
	defer bd.Close()
	if err != nil {
		fmt.Println("error de acceso")
	} else {
		fmt.Println("Acceso satisfactorio")
	}
	//  fmt.Printf("Primer error: ", err)

	err2 := bd.Ping()
	if err2 != nil {
		fmt.Println("error ping")
	} else {
		fmt.Println("Acceso ping")
	}
	//fmt.Printf("Segundo error:", err)

	/*
	   tabla, err2 := bd.Query("SELEC * FROM public.users")
	   if err2 !=nil {
	   	fmt.Printf("error en consulta de tabla\n")
	   }

	   for tabla.Next() {
	   	var name, password string

	   	error3:=tabla.Scan(&name,&password)
	   	if error3!=nil {
	   		fmt.Printf("error recorriendo la tabla\n")
	   	}
	   	fmt.Println(name + " " + password)
	   }
	*/
	bd.Close()
}
