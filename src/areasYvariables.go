package main

import "fmt"

func main() {

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Área del Cuadrado:", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multipliocación
	result = y * x
	fmt.Println("Multiplicación:", result)

	// División
	result = y / x
	fmt.Println("División:", result)

	// Módulo
	result = y % x
	fmt.Println("Módulo:", result)

	// Incrementar
	x++
	fmt.Println("Incrementar var 'x':", x)

	// Decrementar
	x--
	fmt.Println("Incrementar var 'x':", x)

	// Complejo
	a := 10 + 8i
	b := 5 - 10i
	fmt.Println("Número Complejo:", a*b)

	// Área de Rectángulo
	largo := 12
	ancho := 15
	fmt.Println("Área Rectángulo:", largo*ancho)

	// Área de Trapecio
	base_a := 12
	base_b := 15
	altura := 2
	fmt.Println("Área Trapecio:", ((base_a+base_b)/2)*altura)

	// Área de Circulo
	const pi float64 = 3.14
	var radio float64 = 6
	fmt.Println("Área Círculo:", pi*(radio*radio))
}
