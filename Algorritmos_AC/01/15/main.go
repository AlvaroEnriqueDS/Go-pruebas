package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var precio, vender , porcentaje float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, ingrese el precio del articulo"
	porcentaje = 16.5/100
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&precio)
	//SALIDA
	vender = (porcentaje*precio) + precio
	fmt.Println("El precio al cual debe venderlo es de: ", vender)
}


