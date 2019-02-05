package main
//importando paquetes necesarios
import (
"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var numero1 , numero2, burbuja float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese 2 numeros"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&numero1)
	fmt.Scanln(&numero2)
	//SALIDA
	burbuja = numero1
	numero1 = numero2
	numero2 = burbuja
	fmt.Println("se ha cambiado los valores de los numeros: ",numero1,numero2)
}
