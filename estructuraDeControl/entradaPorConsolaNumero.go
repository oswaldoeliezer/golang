/*
Nombre: entraPorConsolaNumero.go
Descripción: Toma dos datos int ingresados por el usuario
e imprime la suma. Para realizar la suma se llamada a una función interna
Autor: @oswaldom876
Fecha:15-01-2017
*/

package main

//Área para importar librerias
import "fmt"

func main() {//Inicio de la función main (bloque de código principal)
        var a int
        var b int

        fmt.Println("Introduce un numero")
        fmt.Scanln(&a)
        fmt.Println("Introduce otro numero")
        fmt.Scanln(&b)
        fmt.S

        //imprime mensaje con formato personalizado
        fmt.Printf("La suma de %d + %d = %d \n", a, b, suma(a,b))
      }//fin de la función main


<<<<<<< HEAD
=======
<<<<<<< HEAD
>>>>>>> refs/remotes/origin/master
      //Descripción de los parametros:
      // func nombreFuncion (var_recibida1 tipo_var_recibida1,
      //              var_recibida2 tipo_var_recibida2) tipo_dato_devuelto
      //              {bloque_de_codigo_a_ejecutar}
<<<<<<< HEAD
=======
=======

      /*Descripción de los parametros:
       func nombreFuncion (var_recibida1 tipo_var_recibida1,
                    var_recibida2 tipo_var_recibida2) tipo_dato_devuelto{

                    bloque_de_codigo_a_ejecutar
                    return (valor a devolver)
                  }
      */
             
>>>>>>> b7e9a9eb7d7c5cdb218fea7f2780f6d71e8e7627
>>>>>>> refs/remotes/origin/master
      func suma(a int,b int) int {//Inicio función suma
               return (a+b)
             }//fin función suma
