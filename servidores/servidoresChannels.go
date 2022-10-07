/*
Consulta de servidores con Channels
aplicando manejo de concurrencias

La ultma vez que tomé el tiempo demora unos
2.22 segundos en verificar todos los servidores
contra los 5.8 de la versión secuencial

Con solo la palabra 'go' no alcanza, es necesario
tener los channels para saber que pasó con esas
subrutinas
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

	// Creación de canales Channels
	canal := make(chan string)

	servidores := []string{
		"http://facebook.com",
		"http://instagram.com",
		"http://wikipedia.com",
		"http://tiktok.com",
	}

	// Versión 1
	// for _, servidor := range servidores {
	// 	go revisarServidor(servidor, canal)
	// }

	// Versión 1
	// Alguien tiene que estar pendiente escuchando este canal
	// y como tego 4 servidores, consultla respuesta de las 4
	// for i := 0; i < len(servidores); i++ {
	// 	fmt.Println(<-canal)
	// }

	// Verisón 2
	// Ahora sustituimos el antiguo for que escuchaba los 4 servidores
	// ahora los construimos para que corra de manera indeterminada
	for {
		for _, servidor := range servidores {
			go revisarServidor(servidor, canal)
		}
		// time.Sleep(time.Second * 1)
		fmt.Println(<-canal)
	}

	// Comparo con el timepo de inicio
	// para saber cuanto demoró en ejecutar
	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de ejecución %s\n", tiempoPaso)
}

// Reviso de los servidores, avisando
// si su funcionamiento es correcto o no responde
func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)

	if err != nil {
		// fmt.Println(servidor, "no esta disponible")
		canal <- servidor + " no esta disponible"
	} else {
		// fmt.Println(servidor, "está funcionando normalmente")
		canal <- servidor + " está funcionando normalmente"
	}
}
