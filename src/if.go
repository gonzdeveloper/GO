package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	valor1 := 1
	valor2 := 1

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	if valor1 == 1 && valor2 == 1 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Es falso")
	}

	//Convertir texto a número
	value, err := strconv.Atoi("53")
	if err != nil {
		fmt.Println("value: ", value)
	} else {
		fmt.Println("value: ", value)
	}

	value, err = strconv.Atoi("asda")
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Println("value: ", value)
	}
}
