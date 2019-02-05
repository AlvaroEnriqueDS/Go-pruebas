package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var opcion int8
	var numero1, numero2  float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, a su calculadora\n" +
			"1.suma\n" +
			"2.resta\n" +
			"3.multiplicacion\n" +
			"4.division"
	//IMPRESION Y LLENADO
		fmt.Println(menu)
		fmt.Scanln(&opcion)
		fmt.Println("inserte los numeros a operar")
		fmt.Scanln(&numero1)
		fmt.Scanln(&numero2)
	//SALIDA
	switch opcion {
	case 1:
		fmt.Println("la suma es: ", numero1+numero2)
	case 2:
		fmt.Println("la resta es: ", numero1-numero2)
	case 3:
		fmt.Println("la multiplicacion es: ", numero1*numero2)
	case 4:
		fmt.Println("la division es: ", numero1/numero2)
	}
}


