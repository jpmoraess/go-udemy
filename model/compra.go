package model

import (
	"errors"
	"time"
)

type Compra struct {
	Data    time.Time
	Mercado string
	Itens   []ItemCompra
	Total   float64
}

type ItemCompra struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func NewCompra(data time.Time, mercado string, itens []ItemCompra) (*Compra, error) {
	if mercado == "" {
		return nil, errors.New("mercado é obrigatório")
	}
	if len(itens) == 0 {
		return nil, errors.New("itens são obrigatórios")
	}
	var total float64
	for _, item := range itens {
		total += item.Preco * float64(item.Quantidade)
	}
	return &Compra{
		Data:    data,
		Mercado: mercado,
		Itens:   itens,
		Total:   total,
	}, nil
}
