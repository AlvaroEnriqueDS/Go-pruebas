package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var numero float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese un numero real"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&numero)
	//SALIDA
	valor := int8(numero)
	fmt.Println("Su valor absoluto es: ",valor)
}


