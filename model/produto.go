package model

import (
	"fmt"
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

		err = data.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			log.Fatal(err.Error())
		}

		p := Produto{
			id, nome, descricao, preco, quantidade,
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

func CreateNewProduct(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectionDatabasePostgres()
	defer db.Close()

	data, err := db.Prepare(`INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4);`)

	if err != nil {
		fmt.Println(err.Error())
	}

	data.Exec(nome, descricao, preco, quantidade)
}
