package main

import "fmt"

func DisplayProducts(products []Product) {

	fmt.Println("\n========== INVENTORY ==========")

	for _, product := range products {
		fmt.Printf(
			"ID: %d | Name: %s | Price: %.2f€ | Stock: %d\n",
			product.ID,
			product.Name,
			product.Price,
			product.Quantity,
		)
	}
}

func AddProduct(products *[]Product, newProduct Product) bool {
	for _, product := range *products {
		if product.ID == newProduct.ID {
			return false
		}
	}

	*products = append(*products, newProduct)
	return true
}

func FindProduct(products []Product, id int) (Product, bool) {
	for _, product := range products {
		if product.ID == id {
			return product, true
		}

	}
	return Product{}, false
}

func UpdateStock(products *[]Product, id int, amount int) bool {
	for i := range *products {
		if (*products)[i].ID == id {

			if (*products)[i].Quantity+amount < 0 {
				return false
			}

			(*products)[i].Quantity += amount
			return true
		}
	}

	return false
}

func DeleteProduct(products *[]Product, id int) bool {
	for i := range *products {
		if (*products)[i].ID == id {

			*products = append((*products)[:i], (*products)[i+1:]...)
			return true
		}

	}
	return false
}
