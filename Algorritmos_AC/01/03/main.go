package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var numero, pulga, salida float32
	//LLENADO DE VARIABLES
	pulga = 2.54
	menu :=
		"Bienvenido, ingrese las pulgadas?"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&numero)
	//SALIDA
	salida = numero*pulga
	fmt.Println("se ha convertido a: ",salida,"  centimatros")
}
