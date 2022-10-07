package main

import (
	"fmt"
	cp "main/GO/src/myPackage"
	"sync"
	"time"
)

type car struct {
	brand string
	year  int
}

func say(text string, wg *sync.WaitGroup) {

	// Para decile que libere esa goroutine
	defer wg.Done()

	fmt.Println(text)
}

func main13() {
	myCar := car{brand: "Ford", year: 2020}

	fmt.Println(myCar)

	// Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"

	fmt.Println(otherCar)

	// Acceso a un struct público de un paquete
	// cp es un alias
	var myCar2 cp.CarPublic
	myCar2.Barnd = "Hiunday"
	myCar2.Year = 2000

	fmt.Println(myCar2)

	cp.Mensaje()

	// Manejo de punteros
	a := 50 //valor
	b := &a // memoria
	// & apunta  a la dirección de memoria
	// * apunta al valor en esa dirección de memoria
	fmt.Println(a, b, *b)

	*b = 101
	fmt.Println(a, b, *b)

	myPC := cp.PC{Ram: 8, Disk: 250, Brand: "Mac"}
	fmt.Println(myPC)
	// Llamo a la función de ese struct
	myPC.Ping()

	myPC.Ram = 16
	fmt.Println(myPC)
	myPC.AumentarRam()
	fmt.Println(myPC)
	myPC.AumentarRam()
	fmt.Println(myPC)

	// Manejo de Interfaces
	myCuadrado := cp.Cuadrado{Base: 2}
	myRectangulo := cp.Rectangulo{Base: 2, Altura: 4}

	cp.Calcular(myCuadrado)
	cp.Calcular(myRectangulo)

	// Lista de Intefaces
	myInterface := []interface{}{"hola", 12, 4.98}
	fmt.Println(myInterface...)

	// Iterar sobre los distintos tipos de datos de ese slice
	for _, v := range myInterface {
		switch v.(type) {
		case int:
			fmt.Println("Es int")
		case string:
			fmt.Println("Es string")
		case float64:
			fmt.Println("Es float64")
		}
	}

	// Manejo de Goroutines
	var wg sync.WaitGroup
	fmt.Println("hello")
	wg.Add(1)
	go say("World", &wg) // Lo ejecuta de forma concurrente
	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios!")

	time.Sleep(time.Second * 1) // NO ES EFICIENTE
}
