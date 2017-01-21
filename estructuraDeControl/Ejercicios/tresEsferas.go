/*
Nombre: tresEsferas.go
Descripción: código para dar respuesta al siguiente postulado:
Supongamos que colocamos dentro de un sombrero tres esferas,
una amarilla, otra azul y otra roja.
Vamos a sacar del sombrero una de esas esferas.
Escriba un código que  imprima el color de esa esfera.
Autor: @oswaldom876
Fecha:11-01-2017
*/

package main

import ( //Área para importar librerias en bloque
    "fmt"
)

func main() { //Inicio de la función main (bloque de código principal)

  //Declaramos las variables
  var esfera1, esfera2, esfera3 = "Amarilla", "Azul", "Roja"

  var esferaSacada ="Amarilla"


//Con un If anidado resolvemos esto
  if esferaSacada == esfera1 {
    fmt.Println("La esfera es " + esfera1)
        } else if esferaSacada == esfera2{
            fmt.Println("La esfera es " + esfera2)
            } else {fmt.Println("La esfera es " + esfera3)}



}//fin de la función main


/*
Salida:
La esfera es Amarilla

*/
