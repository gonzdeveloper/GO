package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

// creamos una variable dentro del tipo task
func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {

	t := task{
		nombre:      "Completar curso",
		descripcion: "Terminar en esta semana",
	}

	fmt.Printf("%+v\n", t)

	t.marcarCompleta()
	t.actualizarNombre("Terminó el curso")
	t.actualizarDescripcion("Completé el curso")

	fmt.Printf("%+v\n", t)

}
