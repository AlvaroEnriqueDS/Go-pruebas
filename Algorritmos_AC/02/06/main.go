package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var unidades, obsequio int8
	var descuento, precio  float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, cuantas unidades ha comrpado"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&unidades)
	fmt.Println("ingrese el precio del producto")
	fmt.Scanln(&precio)
	precio = precio * float32(unidades)

	//SALIDA
	obsequio = unidades/12
	if unidades <=36 {
		descuento = precio * 0.1
	}else {
		descuento = precio * 0.15
	}
	fmt.Println("el monto de la compra es: ",precio)
	fmt.Println("el descuento es: ", descuento)
	fmt.Println("unidades de obsequio: ",obsequio)

}


