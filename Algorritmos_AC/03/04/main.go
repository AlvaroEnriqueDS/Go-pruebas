package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var numero, valor float32
	var i int8
	i = 0
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese numeros"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	for{
		fmt.Scanln(&numero)
		if numero == 0 {
			break
		}
		valor += numero
		i++
	}
	valor_f := valor/float32(i)
	//SALIDA
	fmt.Println("El promedio de los números es: ",valor_f)
}