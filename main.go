package main

import "fmt"

func main() {

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

		var choice int

		fmt.Print("\nChoose an option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			DisplayProducts(products)

		case 2:

			var id int
			var name string
			var price float64
			var quantity int

			fmt.Print("Enter product ID: ")
			fmt.Scan(&id)

			fmt.Print("Enter product name (one word for now): ")
			fmt.Scan(&name)

			fmt.Print("Enter product price: ")
			fmt.Scan(&price)

			fmt.Print("Enter product quantity: ")
			fmt.Scan(&quantity)

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
			fmt.Println("Find Product - Coming soon!")

		case 4:
			fmt.Println("Update Stock - Coming soon!")

		case 5:
			fmt.Println("Delete Product - Coming soon!")

		case 6:
			fmt.Println("Thank you for using GoStock!")
			return

		default:
			fmt.Println("Invalid option. Please choose 1-6.")
		}
	}
}
