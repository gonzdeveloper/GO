package main

import "fmt"

// Creación de la Interface
type animal interface {
	mover() string
}

func moverAnimal(a animal) {
	fmt.Println(a.mover(), "\n")
}

// Creación de la Interface

type perro struct{}

type pez struct{}

type pajaro struct{}

func (perro) mover() string {
	return "Soy un perro y estoy camiando"
}

func (pez) mover() string {
	return "Soy un pez y estoy nadando"
}

func (pajaro) mover() string {
	return "Soy un pájaro y estoy volando"
}

func main() {

	p := perro{}
	moverAnimal(p)
	pe := pez{}
	moverAnimal(pe)
	pa := pajaro{}
	moverAnimal(pa)

}
