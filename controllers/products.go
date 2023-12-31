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
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Erro ao converter ID:", err)
	}
	models.Delete(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	/** Renderiza página de edição */
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Erro ao converter ID:", err)
	}

	productToUpdate := models.FindById(id)
	myTemplate.ExecuteTemplate(w, "Edit", productToUpdate)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println("Erro ao converter ID para inteiro:", err)
		}

		name := r.FormValue("name")
		description := r.FormValue("description")

		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("Erro ao converter preço para float:", err)
		}

		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			log.Println("Erro ao converter quantidade para inteiro:", err)
		}

		models.Update(id, name, description, quantity, price)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
