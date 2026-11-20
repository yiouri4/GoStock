package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle browser preflight requests.
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	switch r.Method {

	case http.MethodGet:
		products, err := LoadProducts()
		if err != nil {
			http.Error(w, "Could not load products", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)

	case http.MethodPost:
		var newProduct Product

		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			http.Error(w, "Invalid product data", http.StatusBadRequest)
			return
		}

		newProduct.Name = strings.TrimSpace(newProduct.Name)

		if newProduct.ID <= 0 || newProduct.Name == "" ||
			newProduct.Price < 0 || newProduct.Quantity < 0 {
			http.Error(w, "Invalid product details", http.StatusBadRequest)
			return
		}

		products, err := LoadProducts()
		if err != nil {
			http.Error(w, "Could not load products", http.StatusInternalServerError)
			return
		}

		added := AddProduct(&products, newProduct)
		if !added {
			http.Error(w, "Product ID already exists", http.StatusConflict)
			return
		}

		err = SaveProducts(products)
		if err != nil {
			http.Error(w, "Could not save products", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newProduct)

	case http.MethodPut:
		var updatedProduct Product

		err := json.NewDecoder(r.Body).Decode(&updatedProduct)
		if err != nil {
			http.Error(w, "Invalid product data", http.StatusBadRequest)
			return
		}

		updatedProduct.Name = strings.TrimSpace(updatedProduct.Name)

		if updatedProduct.ID <= 0 || updatedProduct.Name == "" ||
			updatedProduct.Price < 0 || updatedProduct.Quantity < 0 {
			http.Error(w, "Invalid product details", http.StatusBadRequest)
			return
		}

		products, err := LoadProducts()
		if err != nil {
			http.Error(w, "Could not load products", http.StatusInternalServerError)
			return
		}

		found := false

		for i := range products {
			if products[i].ID == updatedProduct.ID {
				products[i] = updatedProduct
				found = true
				break
			}
		}

		if !found {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		err = SaveProducts(products)
		if err != nil {
			http.Error(w, "Could not save products", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedProduct)

	case http.MethodDelete:
		id := r.URL.Query().Get("id")

		productID, err := strconv.Atoi(id)
		if err != nil || productID <= 0 {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		products, err := LoadProducts()
		if err != nil {
			http.Error(w, "Could not load products", http.StatusInternalServerError)
			return
		}

		deleted := DeleteProduct(&products, productID)
		if !deleted {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		err = SaveProducts(products)
		if err != nil {
			http.Error(w, "Could not save products", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
