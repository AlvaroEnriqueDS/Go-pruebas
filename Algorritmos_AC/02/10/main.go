package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var nota float32
	var pr int8
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, su nota del parcial y seguido cuantos problemas resolvio"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&nota)
	fmt.Scanln(&pr)
	//SALIDA
	if nota >= 10.5 && pr >= 15 {
		fmt.Println("Este alumno está aprobado")
	}else {
		fmt.Println("Este alumno desprobó")
	}
}


