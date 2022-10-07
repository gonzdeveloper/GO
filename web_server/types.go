package main

import (
	"net/http"
)

// El objetivo de esto es tener un hanlder que vaya a ser
// verificado antes, uno puede devolver otro y así sucesivamente
// ... son Handlers que devuelven otros Handlers evaluando
// cierta lógica
type Middleware func(http.HandlerFunc) http.HandlerFunc
