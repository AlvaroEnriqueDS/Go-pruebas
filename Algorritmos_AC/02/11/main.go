package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var num1, num2, num3, burbuja float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese 3 numeros"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&num1)
		fmt.Scanln(&num2)
		fmt.Scanln(&num3)
		burbuja=num1
	//SALIDA
	if burbuja < num2 {
		num1 = num2
		num2 = burbuja

	}else if num2 < num3 {
		burbuja = num2
		num2 = num3
		num3 = burbuja
	}
	fmt.Println("El mayor es: ",num1)
	fmt.Println("El del medio es: ",num2)
	fmt.Println("El menor es: ",num3)

}


