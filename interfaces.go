package main

import (
	"fmt"
	"math"
	"errors"
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

func ExibeError(err error) {
	fmt.Println(err.Error())
}

type ProblemaDeNetwork struct {
	Rede bool
	Hardware bool
}

func (p ProblemaDeNetwork) Error() string {
	if p.Rede {
		return "problema de rede"
	} else if p.Hardware {
		return "problema de hardware"
	} else {
		return "outro problema"
	}
}

func main() {
	ExibeError(errors.New("error"))

	p := ProblemaDeNetwork{
		Rede: true,
		Hardware: false,
	}

	ExibeError(p)
}