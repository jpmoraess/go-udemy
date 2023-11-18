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

	fmt.Println(endereco)
	endereco.Numero = "12302"
	fmt.Println(endereco.Numero)
}
