package main
//importando paquetes necesarios
import (
	"fmt"
)
func main()  {
	//DECLARAR VARIABLES
	var a,b,c,d,e,f, x, y float32
	//LLENADO DE VARIABLES
	menu :=
		"Bienvenido, los numeros a, b ,c ,d ,e y f de las ecuaciones\n" +
			"x = ce-bf/ae-bd  y  y = af-cd/ae-bd"
	//IMPRESION Y LLENADO
	fmt.Println(menu)
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Scanln(&d)
	fmt.Scanln(&e)
	fmt.Scanln(&f)
	//SALIDA
	x = (c*e-b*f)/(a*e-b*d)
	y = (a*f-c*d)/(a*e-b*d)
	fmt.Println("La ecuacion X: ",x)
	fmt.Println("La ecuacion Y: ",y)

}

