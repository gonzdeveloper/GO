package main

import "fmt"

func main() {

	m:= make(map[string]int)

	m["jose"] = 14
	m["pepito"] = 20

	fmt.Println(m)

	// Recorrer map
	for i, v := range m {
		fmt.Println(i,v)
	}

	// Encontrar Valor
	// el ok me devuelve si hubo considencia o no
	value, ok := m["josep"]
	fmt.Println(value, ok)

}