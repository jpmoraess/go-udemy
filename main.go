package main

import (
	"fmt"
	"go-udemy/model"
)

func main() {
	/*
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
		pessoa.CalculaIdade()
		fmt.Println(pessoa.Idade)
	*/

	automovelMoto := model.Automovel{
		Ano:    2022,
		Placa:  "DXO-2020",
		Modelo: "CG",
		Cor:    "Azul",
	}
	moto := model.Moto{
		Automovel:   automovelMoto,
		Cilindradas: 150,
	}
	fmt.Println(moto)
	moto.AlterarAno(2023)
	fmt.Println(moto.Modelo)
	fmt.Println(moto)
}
