package main

import "fmt"

// Otras utilidades https://pkg.go.dev/fmt

func main11() {

	// Declaración de Variables
	helloMessage := "Hello"
	worldWorld := "world"

	// Println
	fmt.Println(helloMessage, worldWorld)
	fmt.Println(helloMessage, worldWorld)

	// Printf
	nombre := "Gonzalo"
	var edad uint8 = 35
	fmt.Printf("%s tiene %d años.\n", nombre, edad)
	// se usa solo 'v' cuando no sabemos que tipo de dato va a venir
	fmt.Printf("%v tiene %d años.\n", nombre, edad)

	// Sprintf
	message := fmt.Sprintf("%s tiene %d años.\n", nombre, edad)
	fmt.Println(message)

	// Tipo de dato de una Variable
	fmt.Printf("helloMessage: %T - ", helloMessage)
	fmt.Printf("edad: %T\n", edad)

}
