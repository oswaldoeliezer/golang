/*
Nombre: lecturaPorConsola.go
Descripción: código que muestra como tomar un dato ingresado por la consola
Autor: @oswaldom876
Fecha:14-01-2017
*/
package main

import (//Área para importar librerias en bloque
	"bufio"
	"fmt"
	"os"
)

func main() {//Inicio de la función main (bloque de código principal)

		 // asigno memoria para una variable de tipo scanner
		 scanner := bufio.NewScanner(os.Stdin)

     fmt.Print("Ingresa un texto")//mensaje a mostrar al usuario

		 scanner.Scan()//leo la entrada por consola

		 //scanner.Text() devuelve lo escrito en la consola como un string
		 texto   := scanner.Text()

		 //Muestro por consola el string ingresado previamente por el usuario
     fmt.Println("El texto ingresado :", texto)

}//fin de la función main
