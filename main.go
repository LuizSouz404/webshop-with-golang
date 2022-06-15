package main

import (
	"html/template"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectionDatabasePostgres() *sql.DB {
	connStr := "user=postgres dbname=postgres password=alura host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := ConnectionDatabasePostgres()
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

	temp.ExecuteTemplate(w, "Index", produtos)
}
