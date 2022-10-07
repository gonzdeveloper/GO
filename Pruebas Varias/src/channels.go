package main

import "fmt"

// Manejo de channels
// Es una forma de organizar las goroutines

// Es bueno indicar si el canal es de entrada chan<-
// o si es de salida <-chan
func mensaje(text string, c chan<- string) {
	c <- text
}

func main() {
	// el uno indica la cantidad de goroutines que va a aceptar
	// buena práctica es ponerle una cantidad
	// si no se le pone la cantidad toma una cantidad dinámica
	c := make(chan string, 1)

	fmt.Println("Hello")

	go mensaje("Bye", c)

	fmt.Println(<-c)
}
