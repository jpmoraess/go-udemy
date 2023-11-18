package model

import "time"

type Pessoa struct {
	Nome           string
	Endereco       Endereco
	DataNascimento time.Time
}

func (p Pessoa) IdadeAtual() int { // Método
	anoNascimento := p.DataNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoNascimento
}

func CalculaIdade(p Pessoa) int {
	anoNascimento := p.DataNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoNascimento
}
