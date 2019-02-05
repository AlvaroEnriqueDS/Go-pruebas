package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var radio, S float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese el RADIO de un circulo"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&radio)
	//SALIDA
	S = 3.14*(radio*radio)
	fmt.Println("los resultados son: \n Superficie: ",S)
}

