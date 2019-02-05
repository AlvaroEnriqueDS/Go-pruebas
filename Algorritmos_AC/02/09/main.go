package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var temperatura, presion  float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese la temperatura"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&temperatura)
	fmt.Println("ingrese la presion")
	fmt.Scanln(&presion)


	//SALIDA
	if temperatura >= 150 {
		fmt.Println("ALERTA la temperatura supera los 150")
	}
	if presion <= 300{
		fmt.Println("Alerta la Presion supera los 300")
	}
	if presion <=300 && temperatura >= 150 {
		fmt.Println("Alerta ambas vriables ya superaron el limite")
	}
	if temperatura <150 && presion >300{
		fmt.Println("Todo está normal")
	}
}


