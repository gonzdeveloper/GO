package main

import "fmt"

func main4() {

	// For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------------")

	// For While
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
