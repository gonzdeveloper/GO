package main

import "fmt"

func main() {
	modulo := 5 % 2

	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("No es par")
	}

	// Otra manera

	switch modulo2 := 10 % 2; modulo2 {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("No es par")
	}

	// Otra forma

	value := 200
	switch {
	case value > 100:
		fmt.Println("Mayor que 100")
	case value < 0:
		fmt.Println("Menor que 0")
	default:
		fmt.Println("No cumple la condición")
	}
}
