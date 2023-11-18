package model

import "fmt"

type Automovel struct {
	Ano    int
	Placa  string
	Modelo string
	Cor    string
}

type Moto struct {
	Automovel
	Cilindradas int
}

type Carro struct {
	Automovel
	Portas            int
	Potencia          int
	DirecaoHidraulica bool
	VidroEletrico     bool
}

func (a *Automovel) AlterarAno(ano int) {
	a.Ano = ano
	fmt.Println("Ano alterado")
}
