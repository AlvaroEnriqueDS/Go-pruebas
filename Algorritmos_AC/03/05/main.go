package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var numero1, numero2, valor_f float32
	numero2 = 100
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese numeros(cierra con 0)"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	for{
		fmt.Scanln(&numero2)
		fmt.Scanln(numero1)
		if numero1 < numero2 {
			valor_f = numero1
		}
		if numero1 == 0 || numero2 == 0 {
			break
		}
	}
	//SALIDA
	fmt.Println("El numero menor es: ",valor_f)
}