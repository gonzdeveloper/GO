package main

import "fmt"

func main6() {

	// Declaración de Constantes
	const pi float64 = 3.14
	const pi2 float64 = 3.14

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de Variables Enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values -> Estos son los valores pode defecto a variables no inicializadas
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Área de un triángulo
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Área cuadrado:", areaCuadrado)

}
