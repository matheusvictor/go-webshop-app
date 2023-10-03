package main

import (
	"html/template"
	"net/http"

	models "alura/go-webshop-app/models/products"

	_ "github.com/lib/pq"
)

var myTemplate = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.FindAll()
	myTemplate.ExecuteTemplate(w, "Index", allProducts)
}
