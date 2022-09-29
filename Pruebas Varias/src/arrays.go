package main

import "fmt"

func main2() {

	// Arrays
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slice
	// no sabemos cuantos elementos vamos a tener
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos de slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
