/*
Consulta de servidores de manera secuencial
(o sea, sin aplicar manejo de concurrencias)

La ultma vez que tomé el tiempo demora unos
5.8 segundos en verificar todos los servidores
*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// Para saber cuando empezó a ejecutar mi programa
	inicio := time.Now()

	servidores := []string{
		"http://facebook.com",
		"http://instagram.com",
		"http://wikipedia.com",
		"http://tiktok.com",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor)
	}

	// Comparo con el timepo de inicio
	// para saber cuanto demoró en ejecutar
	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de ejecución %s\n", tiempoPaso)
}

// Reviso de los servidores, avisando
// si su funcionamiento es correcto o no responde
func revisarServidor(servidor string) {
	_, err := http.Get(servidor)

	if err != nil {
		fmt.Println(servidor, "no esta disponible")
	} else {
		fmt.Println(servidor, "está funcionando normalmente")
	}
}
