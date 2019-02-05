package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var presupuesto float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese el presupuesto de la clinica"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&presupuesto)
	//SALIDA
	cirugia := 0.27*presupuesto
	radiologia := 0.35*presupuesto
	ginecologia := 0.38*presupuesto
	fmt.Println("Cirugia: ",cirugia)
	fmt.Println("radiologia: ",radiologia)
	fmt.Println("ginecologia: ",ginecologia)
}

