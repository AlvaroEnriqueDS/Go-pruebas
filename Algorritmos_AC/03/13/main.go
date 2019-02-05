package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	//var numero int8
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese numeros"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	//fmt.Scanln(&numero)
	for i:=500; i<700; i++{
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}