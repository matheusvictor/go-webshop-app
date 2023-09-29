package main

import (
	"alura/go-webshop-app/produtos"
	"html/template"
	"net/http"
)

var myTemplate = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []produtos.Produto{
		{Nome: "Caneta azul", Descricao: "Azul caneta", Preco: 2, Quantidade: 100},
		{"Tênis", "Air Jordan", 2600, 10},
	}

	myTemplate.ExecuteTemplate(w, "Index", produtos)
}
