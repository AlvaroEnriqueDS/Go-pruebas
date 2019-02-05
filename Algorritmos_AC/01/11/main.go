package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var salario, S float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese su salario"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&salario)
	//SALIDA
	S = 1.19*(salario)
	fmt.Println("Su salario se ha aumentado en un 19%: ",S)
}

