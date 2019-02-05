package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var presion , vol,  temperatura float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese la presion, el volumen y la temperatura"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&presion)
	fmt.Scanln(&vol)
	fmt.Scanln(&temperatura)
	//SALIDA
	masa := (presion*vol)/(0.37*(temperatura+460))
	fmt.Println("La masa de este cuerpo es :", masa)
}

