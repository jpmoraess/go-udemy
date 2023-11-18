package main

import "fmt"

type endereco struct {
	rua    string
	numero string
	cidade string
}

func main() {
	fmt.Println("Iniciando...")

	endereco := endereco{
		rua:    "Rua",
		numero: "123",
		cidade: "Campinas",
	}

	fmt.Println(endereco)
	endereco.numero = "12302"
	fmt.Println(endereco.numero)
}
