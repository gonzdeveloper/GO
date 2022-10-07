package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			// solo para jugar y probar el chqueo
			flag := true

			fmt.Println("Chequeando Autenticación!")

			// si es true, nuestro middleware fue exitoso
			// y tenemos que seguir con otro, si es necesario
			if flag {
				f(w, r)
			} else {
				fmt.Println("No pasó el chequeo!")
				return
			}
		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()

			// Va a retrasar la ejecución de esta función
			// hasta el final de la misma
			defer func() {
				// la diferencia entre log y fmt es que va a poner
				// la hora de la maquina la url estamos llamando
				// y cuanto demora dicha llamada
				log.Println(r.URL.Path, time.Since(start))
			}()
			// y para llamar al siguiente Middleware
			f(w, r)
		}
	}
}
