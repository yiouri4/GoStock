package main

import (
	"encoding/json"
	"os"
)

func LoadProducts() ([]Product, error) {
	data, err := os.ReadFile("products.json")

	if err != nil {
		return nil, err
	}

	var products []Product

	err = json.Unmarshal(data, &products)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func SaveProducts(products []Product) error {
	data, err := json.MarshalIndent(products, "", "    ")

	if err != nil {
		return err
	}

	err = os.WriteFile("products.json", data, 0644)

	if err != nil {
		return err
	}

	return nil
}
