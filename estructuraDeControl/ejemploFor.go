/*
Nombre: ejemploFor.go
Descripción: código que imprime los  números pares contenidos
entre el 1 y el 100
Autor: @oswaldom876
Fecha:13-01-2017
*/

package main

//Área para importar librerias
import "fmt"

func main() { //Inicio de la función main (bloque de código principal)

  fmt.Println("..:: Lista de Números Pares entre el 1 y el 100 ::..\n")

  for i := 1; i <= 100; i++ {
    if  numero  := i; numero % 2  ==  0 {
      fmt.Printf("%d \n", i)
<<<<<<< HEAD
=======

>>>>>>> b7e9a9eb7d7c5cdb218fea7f2780f6d71e8e7627
    }//fin_if

  }//fin_for

}//fin de la función main


/*
Salida:
..:: Lista de Números Pares entre el 1 y el 100 ::..
2
4
6
.
.
.
100

*/
