package routes

import (
	"ecommerce-inventory-service/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/product", handlers.AddProductHandler).Methods(http.MethodPost)
	router.HandleFunc("/product/{id}", handlers.GetProductHandler).Methods(http.MethodGet)
}
