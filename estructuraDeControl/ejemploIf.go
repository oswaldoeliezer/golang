/*
Nombre: ejemploIf.go
Descripción: código que muestra la funcionalidad del If en Go
Autor: @oswaldom876
Fecha:11-01-2017
*/

package main

//Área para importar librerias
import "fmt"

func main() { //Inicio de la función main (bloque de código principal)

  //Declaramos las variables
  var x, y = 15 , 11

  if x > y {
    fmt.Printf("%d es mayor que %d\n", x , y)
  } else {
    fmt.Printf("%d es mayor que %d\n", y , x)
  }

}//fin de la función main


/*
Salida:
15 es mayor que 11

*/
