package model

import (
	"time"
)

type Pessoa struct {
	Nome           string
	Endereco       Endereco
	DataNascimento time.Time
	Idade          int
}

func (p *Pessoa) CalculaIdade() { // Método
	anoNascimento := p.DataNascimento.Year()
	anoAtual := time.Now().Year()
	p.Idade = anoAtual - anoNascimento
}
