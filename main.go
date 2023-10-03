package main

import (
	"net/http"

	"alura/go-webshop-app/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
