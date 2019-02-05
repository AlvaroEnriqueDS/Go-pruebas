package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
		//DECLARAR VARIABLES
		var grados, divi, celcius float32
		//LLENADO DE VARIABLES
		divi = 1.8
		celcius = (grados-32)/divi
		menu :=
		"Bienvenido, ingrese los grados Fahrenheit?"
		//IMPRESION Y LLENADO
		fmt.Println(menu)
		fmt.Scanln(&grados)
		//SALIDA
		fmt.Println("se ha convertido a: ",celcius," grados celcius")
}


