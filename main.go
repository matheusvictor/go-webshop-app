package main

import (
	"html/template"
	"net/http"

	"alura/go-webshop-app/database"
	"alura/go-webshop-app/products"
)

var myTemplate = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := database.GetDBInstance()
	defer db.Close() // defer será executado por último

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []products.Product{
		{Name: "Caneta azul", Description: "Azul caneta", Price: 2, Quantity: 100},
		{"Tênis", "Air Jordan", 2600, 10},
	}

	myTemplate.ExecuteTemplate(w, "Index", products)
}
