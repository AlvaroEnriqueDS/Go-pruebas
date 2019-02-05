package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var radio , altura,  V, S float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese el RADIO y la ALTURA del cilindo"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&radio)
	fmt.Scanln(&altura)
	//SALIDA
	V = (3.14*(radio*radio)*altura)
	S =  (2*3.14)*radio*altura
	fmt.Println("los resultados son\nVolumen: ", V,"\nArea: ",S)
}

