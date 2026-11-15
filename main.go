package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	products, err := LoadProducts()

	if err != nil {
		fmt.Println("Error loading products:", err)
		return
	}

	for {
		fmt.Println("\n========== GoStock ==========")
		fmt.Println("1. Display Products")
		fmt.Println("2. Add Product")
		fmt.Println("3. Find Product")
		fmt.Println("4. Update Stock")
		fmt.Println("5. Delete Product")
		fmt.Println("6. Exit")

		fmt.Print("\nChoose an option: ")

		if !scanner.Scan() {
			return
		}

		choice, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			continue
		}

		switch choice {

		case 1:
			DisplayProducts(products)

		case 2:

			var id int
			var name string
			var price float64
			var quantity int

			fmt.Print("Enter product ID: ")
			if !scanner.Scan() {
				return
			}

			id, err = strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Invalid ID! Please enter a number.")
				break
			}

			fmt.Print("Enter product name: ")
			if !scanner.Scan() {
				return
			}

			name = scanner.Text()

			fmt.Print("Enter product price: ")
			if !scanner.Scan() {
				return
			}

			price, err = strconv.ParseFloat(scanner.Text(), 64)

			if err != nil {
				fmt.Println("Invalid price!")
				break
			}

			fmt.Print("Enter product quantity: ")
			if !scanner.Scan() {
				return
			}

			quantity, err = strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Invalid quantity!")
				break
			}

			if id <= 0 {
				fmt.Println("ID must be greater than 0")
				break
			}

			if name == "" {
				fmt.Println("Name cannot be empty")
				break
			}

			if price < 0 {
				fmt.Println("Price cannot be negative")
				break
			}

			if quantity < 0 {
				fmt.Println("Quantity cannot be negative")
				break
			}

			newProduct := Product{
				ID:       id,
				Name:     name,
				Price:    price,
				Quantity: quantity,
			}

			added := AddProduct(&products, newProduct)

			if !added {
				fmt.Println("Error: A product with this ID already exists!")
				break
			}

			err = SaveProducts(products)

			if err != nil {
				fmt.Println("Error saving product:", err)
				return
			}

			fmt.Println("Product added successfully!")

		case 3:

			var id int

			fmt.Print("Enter product ID: ")
			if !scanner.Scan() {
				return
			}

			id, err = strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Invalid ID!")
				break
			}

			foundProduct, found := FindProduct(products, id)

			if found {
				fmt.Printf(
					"ID: %d | Name: %s | Price: €%.2f | Stock: %d\n",
					foundProduct.ID,
					foundProduct.Name,
					foundProduct.Price,
					foundProduct.Quantity,
				)
			} else {
				fmt.Println("Product not found!")
			}

		case 4:

			var id int
			var amount int

			fmt.Print("Enter product ID: ")
			if !scanner.Scan() {
				return
			}

			id, err = strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Invalid ID!")
				break
			}

			fmt.Print("Enter quantity change (use negative number to decrease): ")
			if !scanner.Scan() {
				return
			}

			amount, err = strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Invalid quantity!")
				break
			}

			updated := UpdateStock(&products, id, amount)

			if !updated {
				fmt.Println("Stock update failed! Product not found or insufficient stock.")
				break
			}

			err = SaveProducts(products)

			if err != nil {
				fmt.Println("Error saving inventory:", err)
				return
			}

			fmt.Println("Stock updated successfully!")

		case 5:

			var id int

			fmt.Print("Enter product ID to delete: ")
			if !scanner.Scan() {
				if err := scanner.Err(); err != nil {
					fmt.Println("Error reading input:", err)
				}

				return
			}

			id, err = strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Invalid ID!")
				break
			}

			deleted := DeleteProduct(&products, id)

			if !deleted {
				fmt.Println("Product not found!")
				break
			}

			err = SaveProducts(products)

			if err != nil {
				fmt.Println("Error saving inventory:", err)
				return
			}

			fmt.Println("Product deleted successfully!")

		case 6:
			fmt.Println("Thank you for using GoStock!")
			return

		default:
			fmt.Println("Invalid option. Please choose 1-6.")
		}
	}
}
