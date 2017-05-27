/*
<<<<<<< HEAD
Nombre: entradaPorConsolaNumero.go
Descripción: Código que permite la suma de dos números enteros positivos
=======
Nombre: entraPorConsolaNumero.go
Descripción: En produccion
>>>>>>> b7e9a9eb7d7c5cdb218fea7f2780f6d71e8e7627
Autor: @oswaldom876
Fecha:15-01-2017
*/

<<<<<<< HEAD
package main
=======
package entraPorConsolaNumero
>>>>>>> b7e9a9eb7d7c5cdb218fea7f2780f6d71e8e7627

//Área para importar librerias
import "fmt"

func main() {//Inicio de la función main (bloque de código principal)
        var a int
        var b int

<<<<<<< HEAD
        fmt.Print("Introduce un numero:")
        fmt.Scanln(&a)
        for  a < 0 {
          //fmt.Scanln(&a)
          fmt.Print("Error, Ingrese un dato numerico:")
          fmt.Scanln(&a)
        }
        //fmt.Scanln(&a)
        fmt.Print("Introduce otro numero:")
        fmt.Scanln(&b)
        for  b < 0 {
          //fmt.Scanln(&a)
          fmt.Print("Error, Ingrese un dato numerico:")
          fmt.Scanln(&b)
        }
=======
        fmt.Println("Introduce un numero")
        fmt.Scanln(&a)
        for  a == 0 {
          fmt.Scanln(&a)
          fmt.Println("Error, Ingrese un dato numerico")
          //fmt.Scanln(&a)
        }
        //fmt.Scanln(&a)
        fmt.Println("Introduce otro numero")
        fmt.Scanln(&b)
>>>>>>> b7e9a9eb7d7c5cdb218fea7f2780f6d71e8e7627

        //imprime mensaje con formato personalizado
        fmt.Printf("La suma de %d + %d = %d \n", a, b, suma(a,b))
      }//fin de la función main


      //Descripción de los parametros:
      // func nombreFuncion (var_recibida1 tipo_var_recibida1,
      //              var_recibida2 tipo_var_recibida2) tipo_dato_devuelto
      //              {bloque_de_codigo_a_ejecutar}
      func suma(a int,b int) int {//Inicio función suma
               return (a+b)
             }//fin función suma
