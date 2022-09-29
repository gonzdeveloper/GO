package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

// Recive function
func (calc) operacion(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {
	case "-":
		return operador1 - operador2
	case "+":
		return operador1 + operador2
	case "*":
		return operador1 * operador2
	case "/":
		return operador1 / operador2
	case "%":
		return operador1 % operador2
	default:
		fmt.Println("Desconocemos la operación!")
		return 0
	}
}

func parsear(entrada string) int {
	operador, err := strconv.Atoi(entrada)

	if err != nil {
		fmt.Println("Verifique la entrad, puede haber caracteres no admitidos!")
		return 0
	} else {
		return operador
	}

}

func leerEntrada() string {
	fmt.Print("Esta Calculadora de ENTEROS solo soporta +, -, /, %\nIntroduce la operación que quieres realizar: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func queOperacionEs(entrada string) string {

	if strings.Contains(entrada, "+") {
		return "+"
	} else {
		if strings.Contains(entrada, "-") {
			return "-"
		} else {
			if strings.Contains(entrada, "/") {
				return "/"
			} else {
				return "*"
			}
		}
	}

}

func main() {

	entrada := leerEntrada()
	operador := queOperacionEs(entrada)
	c := calc{}
	fmt.Println("Resultado:", c.operacion(entrada, operador))

}
