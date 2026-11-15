package main

import "fmt"

func main() {
	fmt.Println("==============================")
	fmt.Println("       WELCOME TO GoStock")
	fmt.Println("==============================")
	fmt.Println("Retail Inventory Management System")
	fmt.Println("Starting application...")

	products := []Product{
		{ID: 1, Name: "Milk", Price: 1.22, Quantity: 50},
		{ID: 2, Name: "Bread", Price: 0.9, Quantity: 40},
		{ID: 3, Name: "Eggs", Price: 2.5, Quantity: 40},
		{ID: 4, Name: "Coffee", Price: 5.15, Quantity: 20},
		{ID: 5, Name: "Cheese", Price: 2.8, Quantity: 30},
	}
	DisplayProducts(products)

	newProduct := Product{
		ID:       6,
		Name:     "Butter",
		Price:    1.34,
		Quantity: 35,
	}

	AddProduct(&products, newProduct)
	DisplayProducts(products)

	foundProduct, found := FindProduct(products, 3)
	if found {
		fmt.Println("\nPRODUCT FOUND!")
		fmt.Println("Name:", foundProduct.Name)
		fmt.Printf("Price: %.2f€\n", foundProduct.Price)
		fmt.Println("Stock:", foundProduct.Quantity)
	} else {
		fmt.Println("\nPRODUCT NOT FOUND!")
	}

	success := UpdateStock(&products, 3, -5)

	if success {
		fmt.Println("Stock updated successfully!")
	} else {
		fmt.Println("Stock update failed!")
	}

	DisplayProducts(products)
}
