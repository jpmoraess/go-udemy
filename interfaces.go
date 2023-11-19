package main

import (
	"fmt"
	"math"
)

type Geometria interface {
	area() float64
}

type Retangulo struct {
	Largura, Altura float64
}

func (r Retangulo) area() float64 {
	return r.Largura * r.Altura
}

type Circulo struct {
	Radius float64
}

func (c Circulo) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func ExibeAreaRetangulo(retangulo Retangulo) {
	area := retangulo.Largura * retangulo.Altura
	fmt.Println("Área do retangulo:", area)
}

func ExibeAreaCirculo(circulo Circulo) {
	area := math.Pi * circulo.Radius * circulo.Radius
	fmt.Println("Área do circulo:", area)
}

func ExibeGeometria(g Geometria) {
	fmt.Println(g.area())
}

func main() {
	fmt.Println("Inicializando...")

	retangulo := Retangulo{
		Largura: 1,
		Altura: 2,
	}

	circulo := Circulo{
		Radius: 3,
	}

	//ExibeAreaRetangulo(retangulo)
	//ExibeAreaCirculo(circulo)

	ExibeGeometria(circulo)
	ExibeGeometria(retangulo)
}
