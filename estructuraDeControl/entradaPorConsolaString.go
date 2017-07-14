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

<<<<<<< HEAD
     fmt.Print("Ingresa un texto")//mensaje a mostrar al usuario
=======
     fmt.Print("Ingresa un texto: ")//mensaje a mostrar al usuario
>>>>>>> b7e9a9eb7d7c5cdb218fea7f2780f6d71e8e7627

		 scanner.Scan()//leo la entrada por consola

		 //scanner.Text() devuelve lo escrito en la consola como un string
		 texto   := scanner.Text()

		 //Muestro por consola el string ingresado previamente por el usuario
<<<<<<< HEAD
     fmt.Println("El texto ingresado :", texto)

}//fin de la función main
=======
     fmt.Println("El texto ingresado: ", texto)

}//fin de la función main

/*
Ingresa un texto: textoIngresado
El texto ingresado :textoIngresado

*/
>>>>>>> b7e9a9eb7d7c5cdb218fea7f2780f6d71e8e7627
