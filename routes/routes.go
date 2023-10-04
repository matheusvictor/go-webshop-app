package routes

import (
	"net/http"

	controllers "alura/go-webshop-app/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}

func CreateProduct() {

}
