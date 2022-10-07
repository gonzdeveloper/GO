package main

import (
	"net/http"
)

/*
Acá vamos a poder crear reglas de ruteo
desde que direcciones nos piden pticiones y a donde
pasamos con cada petición

Cada Router debe tener conocimiento de que handler maneja cada ruta
*/

type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	// Con 'Fprintf' asignamos un escritor para poder enviar
	// nuestro mensaje y se muestra en el navegador
	handler, exist := r.FindHandler(request.URL.Path)

	if !exist {
		// No existe la dirección
		w.WriteHeader(http.StatusNotFound)
		return // Para salir de la función
	}

	handler(w, request)
}

// Para buscar las rutas, si existen o no..?
func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.rules[path]
	return handler, exist
}
