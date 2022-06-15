package controllers

import (
	"html/template"
	"luiz/shop/model"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := model.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", allProducts)
}

func Create(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Create", nil)
}
