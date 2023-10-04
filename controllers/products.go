package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	models "alura/go-webshop-app/models/products"
)

var myTemplate = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.FindAll()
	myTemplate.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		desc := r.FormValue("descricao")

		price, priceError := strconv.ParseFloat(r.FormValue("preco"), 64)
		if priceError != nil {
			log.Println("Erro na conversão do preço:", priceError)
		}

		quant, quantError := strconv.Atoi(r.FormValue("quantidade"))
		if priceError != nil {
			log.Println("Erro na conversão da quantidade:", quantError)
		}

		models.Create(name, desc, price, quant)
	}
	http.Redirect(w, r, "/", 303)
}
