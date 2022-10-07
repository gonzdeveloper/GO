/*
Se encarga de manejar una mejor segmentación de las rutas
Cada Router debe tener conocimiento de que handler maneja cada ruta

Simpre van a tener un escritor w http.ResponseWriter
y un request r *http.Request que es donde se almacena toda
la data que el cliente va enviando
*/
package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde el Handler")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola desde el Endpoint de la API")
}
