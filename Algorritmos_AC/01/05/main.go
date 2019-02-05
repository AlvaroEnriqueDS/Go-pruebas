package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var numero1 , numero2 float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese 2 numeros reales"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&numero1)
	fmt.Scanln(&numero2)
	//SALIDA
	suma := numero1 + numero2
	resta := numero1-numero2
	producto := numero1*numero2
	divi := numero1/numero2
	fmt.Println("se ha calculado su:\n" +
		"suma: ",suma,"\nresta: ",resta,"\nproducto: ", producto,"\ndividision: ", divi)
}

