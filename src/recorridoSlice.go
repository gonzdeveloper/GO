package main

import (
    "fmt"
    "strings"
)



func isPalindromo(text string){
	var textReverse string

	for i:= len(text)-1; i>=0; i-- {
		textReverse += string(text[i])
	}

	if strings.ToLower(text) == strings.ToLower(textReverse) {
		fmt.Println(text, " Es Palíndromo")
	}else{
		fmt.Println(text, " No es Palíndromo")
	}
}


func main(){
	
	slice:= []string{"hola", "como", "estas"}

	// imprimir clave y valor
	for i, valor := range slice{
		fmt.Println(i, valor)
	}

	// imprimir clave
	for i, _ := range slice{
		fmt.Println(i)
	}

	// imprimir clave de otra manera
	for i := range slice{
		fmt.Println(i)
	}

	// imprimir valor
	for _, valor := range slice{
		fmt.Println(valor)
	}


	isPalindromo("amor a Roma")
	isPalindromo("amor a rom")
}