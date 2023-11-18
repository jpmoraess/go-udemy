package main

import (
	"fmt"
	"go-udemy/model"
)

func main() {
	fmt.Println("Iniciando...")

	endereco := model.Endereco{
		Rua:    "Rua",
		Numero: "123",
		Cidade: "Campinas",
	}

	pessoa := model.Pessoa{
		Nome:     "João Pedro",
		Endereco: endereco,
	}

	fmt.Println(pessoa)
	fmt.Println(endereco)
	endereco.Numero = "12302"
	fmt.Println(endereco.Numero)
}
