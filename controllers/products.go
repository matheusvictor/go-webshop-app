package controllers

import (
	"html/template"
	"net/http"

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
