package controllers

import (
	"html/template"
	"log"
	"luiz/shop/model"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := model.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", allProducts)
}

func Create(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Create", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	quantidade := r.FormValue("quantidade")

	precoConvert, err := strconv.ParseFloat(preco, 64)

	if err != nil {
		log.Println("Erro into convert price")
	}

	quantidadeConvert, err := strconv.Atoi(quantidade)

	if err != nil {
		log.Println("Erro into convert quantity")
	}

	model.CreateNewProduct(nome, descricao, precoConvert, quantidadeConvert)

	http.Redirect(w, r, "/", 301)
}
