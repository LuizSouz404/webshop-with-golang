package model

import (
	"log"
	"luiz/shop/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetAllProducts() []Produto {
	db := db.ConnectionDatabasePostgres()
	defer db.Close()

	data, err := db.Query("SELECT * FROM products")

	if err != nil {
		log.Fatal(err)
	}

	// p := Produto{}
	produtos := []Produto{}

	for data.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = data.Scan(&id, &descricao, &nome, &quantidade, &preco)

		if err != nil {
			log.Fatal(err.Error())
		}

		p := Produto{
			id, descricao, nome, preco, quantidade,
		}
		// p.Id=id
		// p.Descricao=descricao
		// p.Nome=nome
		// p.Preco=preco
		// p.Quantidade=quantidade

		produtos = append(produtos, p)
	}

	return produtos
}
