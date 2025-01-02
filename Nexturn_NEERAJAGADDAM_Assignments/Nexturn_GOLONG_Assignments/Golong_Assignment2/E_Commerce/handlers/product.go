package handlers

import (
	"ecommerce-inventory-service/config"
	"ecommerce-inventory-service/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("INSERT INTO products (name, description, price, stock, category_id) VALUES (?, ?, ?, ?, ?)",
		product.Name, product.Description, product.Price, product.Stock, product.CategoryID)

	if err != nil {
		http.Error(w, "Failed to add product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product added successfully"})
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	row := config.DB.QueryRow("SELECT * FROM products WHERE id = ?", id)
	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID)

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}
