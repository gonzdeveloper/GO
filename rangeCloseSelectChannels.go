package main

import (
	"fmt"
)

// Range, Close, Select Channels

func message(text string, c chan string) {
	c <- text
}

func main14() {

	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	// Para ver la cantidad de datos que tengo
	// dentr del Channel y cuál es la capacidad
	fmt.Println(len(c), cap(c))

	// Recorrido de los valores dentro del Channel
	// Range y Close

	// close, cierra el canal luego de usuaralo
	// y no podemos agregar más mensajes
	// es una buena práctica cerrar el channel luego de saber que no va
	// a recibir más datos
	close(c)

	// range, para realizar la recorrida dentro del channel
	// Es importante cerrar el channel antes de realizar la regorrida
	// para avisar al runtime de Go que no espere por nuevos datos
	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)
	go message("mensaje 1", email1)
	go message("mensaje 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email 1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email 2", m2)
		}
	}

}
