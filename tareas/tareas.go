package main

import (
	"fmt"
)

type taskList struct {
	tasks []*task
}

func (tl *taskList) agregarALista(t *task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) eliminarDeLista(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (tl *taskList) imprimirLista() {
	// no me intereza el indice
	for _, tarea := range tl.tasks {
		fmt.Println("Nombre:", tarea.nombre)
		fmt.Println("Descripción:", tarea.descripcion)
	}
}

func (tl *taskList) imprimirListaCompletado() {
	// no me intereza el indice
	for _, tarea := range tl.tasks {
		if tarea.completado {
			fmt.Println("Nombre:", tarea.nombre)
			fmt.Println("Descripción:", tarea.descripcion)
		}
	}
}

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

	t1 := &task{
		nombre:      "Completar curso",
		descripcion: "Terminar en esta semana",
	}

	t2 := &task{
		nombre:      "Completar curso Python",
		descripcion: "Terminar en esta semana Python",
	}

	t3 := &task{
		nombre:      "Completar curso NodeJS",
		descripcion: "Terminar en esta semana NodeJS",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	fmt.Println(lista.tasks[0])

	lista.agregarALista(t3)

	for i := 0; i < len(lista.tasks); i++ {
		fmt.Println("index", i, "nombre", lista.tasks[i].nombre)
	}

	// la forma más utilizada y más elegante
	for index, tarea := range lista.tasks {
		fmt.Println("index", index, "nombre", tarea.nombre)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("con 'break' me voy")
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("con 'continue' me lo salteo y sigo")
			continue
		}
		fmt.Println(i)
	}

	lista.imprimirLista()
	lista.tasks[0].marcarCompleta()
	fmt.Println("Estas son las lista completas:")
	lista.imprimirListaCompletado()

}
