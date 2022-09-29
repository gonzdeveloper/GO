package myPackage

import "fmt"

type Cuadrado struct {
	Base float64
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

func (c Cuadrado) area() float64 {
	return c.Base * c.Base
}

func (r Rectangulo) area() float64 {
	return r.Base * r.Altura
}

type Figuras2D interface {
	area() float64
}

func Calcular(f Figuras2D) {
	fmt.Println("Área:", f.area())
}
