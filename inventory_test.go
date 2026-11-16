package main

import "testing"

func TestAddProduct(t *testing.T) {
	products := []Product{}

	newProduct := Product{ID: 1, Name: "Test Product", Price: 10.0, Quantity: 5}

	result := AddProduct(&products, newProduct)
	if result == false {
		t.Errorf("Expected product to be added, but it was not.")
	}
	if len(products) != 1 {
		t.Errorf("Expected products length to be 1, but got %d", len(products))
	}
}
func TestAddProductDuplicateID(t *testing.T) {
	products := []Product{{ID: 1, Name: "Test Product", Price: 10.0, Quantity: 5}}

	newProduct := Product{ID: 1, Name: "Duplicate Product", Price: 15.0, Quantity: 3}

	result := AddProduct(&products, newProduct)
	if result == true {
		t.Errorf("Expected product not to be added, but it was.")
	}
	if len(products) != 1 {
		t.Errorf("Expected products length to be 1, but got %d", len(products))
	}
}
func TestFindProduct(t *testing.T) {
	products := []Product{{ID: 1, Name: "Test Product", Price: 10.0, Quantity: 5}}

	foundProduct, found := FindProduct(products, 1)
	if !found {
		t.Errorf("Expected product to be found, but it was not.")
	}
	if foundProduct.ID != 1 {
		t.Errorf("Expected product ID to be 1, but got %d", foundProduct.ID)
	}
}
func TestFindProductNotFound(t *testing.T) {
	products := []Product{{ID: 1, Name: "Test Product", Price: 10.0, Quantity: 5}}

	_, found := FindProduct(products, 2)
	if found {
		t.Errorf("Expected product to not be found, but it was.")
	}
}
func TestUpdateStock(t *testing.T) {
	products := []Product{{ID: 1, Name: "Test Product", Price: 10.0, Quantity: 5}}

	result := UpdateStock(&products, 1, 10)
	if result == false {
		t.Errorf("Expected stock to be updated, but it was not.")
	}
	if products[0].Quantity != 15 {
		t.Errorf("Expected product quantity to be 10, but got %d", products[0].Quantity)
	}
}
func TestUpdateStockPreventNegative(t *testing.T) {
	products := []Product{{ID: 1, Name: "Test Product", Price: 10.0, Quantity: 5}}
	result := UpdateStock(&products, 1, -10)
	if result == true {
		t.Errorf("Expected stock update to fail due to negative quantity, but it succeeded.")

	}
	if products[0].Quantity != 5 {
		t.Errorf("Expected product quantity to remain 5, but got %d", products[0].Quantity)
	}
}
func TestDeleteProduct(t *testing.T) {
	products := []Product{{ID: 1, Name: "Test Product", Price: 10.0, Quantity: 5}}
	products = append(products, Product{ID: 2, Name: "Another Product", Price: 15.0, Quantity: 3})
	result := DeleteProduct(&products, 1)
	if result == false {
		t.Errorf("Expected product to be deleted, but it was not.")
	}
	if len(products) != 1 {
		t.Errorf("Expected products length to be 1, but got %d", len(products))
	}
	if products[0].ID != 2 {
		t.Errorf("Expected remaining product ID to be 2, but got %d", products[0].ID)
	}
}
func TestDeleteProductNotFound(t *testing.T) {
	products := []Product{{ID: 1, Name: "Test Product", Price: 10.0, Quantity: 5}}
	result := DeleteProduct(&products, 2)
	if result == true {
		t.Errorf("Expected product deletion to fail, but it succeeded.")
	}
	if len(products) != 1 {
		t.Errorf("Expected products length to remain 1, but got %d", len(products))
	}
	if products[0].ID != 1 {
		t.Errorf("Expected remaining product ID to be 1, but got %d", products[0].ID)
	}
}
func TestLoadProducts(t *testing.T) {
	products, err := LoadProducts()
	if err != nil {
		t.Errorf("Expected products to be loaded, but got error: %v", err)
	}
	if len(products) == 0 {
		t.Fatalf("Expected products to be loaded, but got empty list.")
	}
}
