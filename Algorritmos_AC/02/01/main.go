package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var num1,  num2 float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese 2 numeros"
	//IMPRESION Y LLENADO
		fmt.Println(menu)
		fmt.Scanln(&num1)
		fmt.Scanln(&num2)
	//SALIDA
	if num1 > num2 {
		fmt.Println("el mayor es el primer numero")
	}else if num1 == num2{
		fmt.Println("los numeros sin iguales")
	}else {
		fmt.Println("El mayor es el segundo numero")
	}
}


