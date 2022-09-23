package main

import "fmt"

func main() {

	//Defer
	defer fmt.Println("Chau, cierro base de datos")
	fmt.Println("Hola!")

	// Continue y Break
	for i := 0; i < 10; i++ {

		fmt.Println(i)

		// Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
