package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

// Instanciar nuestro servidor
// '*' para evitar que se estén usando copias
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

// Escuchar conexiones
// '*' para evitar que se estén usando copias
func (s *Server) Listen() error {

	// el punto de inicio de mi web
	http.Handle("/", s.router)

	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		// Tenemos un error
		return err
	} else {
		// Todo salió ok
		return nil
	}
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
}

// Chqueo de Middleware, se agregan los chqueos, como no sabemos
// cuantos middlewares pueden ser se le agregan los tres puntos al argumento
/*
Si por ejemplo en el slice de middlewares hubiera 3 de ellos, seria el equivalente a anidar 3 veces la función:
fFinal = m3(m2(m1(f)))
*/
func (s *Server) AddMiddleware(f http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	for _, m := range middleware {
		f = m(f)
	}
	return f
}
