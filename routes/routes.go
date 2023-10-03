package routes

import (
	"net/http"

	controllers "alura/go-webshop-app/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}