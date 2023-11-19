package main

import (
	"fmt"
	"go-udemy/model"
	"time"
)

func main2() {
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

	/*
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
	*/

	/*
		itens := make([]model.ItemCompra, 0)
		itens = append(itens, model.ItemCompra{
			Nome: "Arroz",
		})
		itens = append(itens, model.ItemCompra{
			Nome: "Feijáo",
		})
		compra := model.Compra{
			Data:    time.Now(),
			Mercado: "Mercado X",
			Itens:   itens,
		}
		fmt.Println(compra)
	*/

	itens := make([]model.ItemCompra, 0)
	itens = append(itens, model.ItemCompra{
		Nome:       "Arroz",
		Preco:      18.0,
		Quantidade: 2,
	})
	itens = append(itens, model.ItemCompra{
		Nome:       "Farofa",
		Preco:      7.0,
		Quantidade: 1,
	})
	compra, err := model.NewCompra(time.Now(), "Mercado Jota", itens)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(*compra)
	}
}
