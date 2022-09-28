package main

import "fmt"

func normalFuncion(message string) {
	fmt.Println(message)
}

func tripleArgumeto(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func dobleReturn(a int) (c, d int) {
	return a, a * 2
}

func main5() {

	normalFuncion("hola mundo")
	tripleArgumeto(1, 2, "hola")

	value := returnValue(4)
	fmt.Println(value)

	value1, value2 := dobleReturn(4)
	fmt.Println(value1, value2)

	// por si quiero solamente obtener un valor y descartar el resto
	value1, _ = dobleReturn(10)
	fmt.Println(value1, value2)
}
