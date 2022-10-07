package main

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
