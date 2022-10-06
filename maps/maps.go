package main

import "fmt"

func main() {
	//estamos diciendo que todas las llaves van a ser de tipo string
	// pero cuando la llamemos nos va a devolver un tipo int
	m1 := make(map[string]int)

	m1["a"] = 8

	fmt.Println(m1)
}
