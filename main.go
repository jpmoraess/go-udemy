package main

import (
	"fmt"
	"go-udemy/model"
	"time"
)

func main() {
	fmt.Println("Iniciando...")

	endereco := model.Endereco{
		Rua:    "Rua",
		Numero: "123",
		Cidade: "Campinas",
	}

	pessoa := model.Pessoa{
		Nome:           "João Pedro",
		Endereco:       endereco,
		DataNascimento: time.Date(1997, 3, 6, 20, 10, 0, 0, time.Local),
	}

	fmt.Println(pessoa)
	fmt.Println(endereco)
	idade := model.CalculaIdade(pessoa)
	idade = pessoa.IdadeAtual()
	fmt.Println(idade)
}
