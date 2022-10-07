/*
Por un tema de problemas con referancias puse todo en el mismo archivo
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	server := NewServer(":3000")

	// Cargo el mapeo de páginas
	server.Handle("/", HandleRoot)

	// Versión 1 sin chequeo
	// server.Handle("/api", HandleHome)

	// Verisón 2 con chequeo de autenticación
	//server.Handle("/api", server.AddMiddleware(HandleHome, CheckAuth()))

	// Versión 3 con ceque de autenticación más la ejecución del logging
	server.Handle("/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))

	// Ahora escucho al servidor y sus peticiones
	server.Listen()
}

// ------------------- Inicio Servidor

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

// ------------------- Servidor

// ------------------- Inicio Router

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

// ------------------- Router

// ------------------- Inicio Handler
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde el Handler")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola desde el Endpoint de la API")
}

// ------------------- Handler

// ------------------- Inicio Types

// El objetivo de esto es tener un hanlder que vaya a ser
// verificado antes, uno puede devolver otro y así sucesivamente
// ... son Handlers que devuelven otros Handlers evaluando
// cierta lógica
type Middleware func(http.HandlerFunc) http.HandlerFunc

// ------------------- Types

// ------------------- Inicio Middleware

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

// ------------------- Middleware
