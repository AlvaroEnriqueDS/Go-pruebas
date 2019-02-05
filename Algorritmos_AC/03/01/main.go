package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var numero, valor float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese 10 numeros"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	for i := 0; i<10; i++ {
		fmt.Scanln(&numero)
		valor = numero +valor
	}
	valor_f := valor/10
	//SALIDA
	fmt.Println("El promedio de los números es: ",valor_f)
}


