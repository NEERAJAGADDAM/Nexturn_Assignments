package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Stock       int
	Category    string
	Description string
}

var inventory []Product

// AddProduct adds a new product to the inventory
func AddProduct(id int, name string, price interface{}, stock int, category string, description string) error {
	convertedPrice, ok := price.(float64)
	if !ok {
		return errors.New("price must be a float64")
	}
	inventory = append(inventory, Product{
		ID:          id,
		Name:        name,
		Price:       convertedPrice,
		Stock:       stock,
		Category:    category,
		Description: description,
	})
	return nil
}

// UpdateStock updates the stock of a product
func UpdateStock(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}
	for i, product := range inventory {
		if product.ID == id {
			inventory[i].Stock = newStock
			return nil
		}
	}
	return errors.New("product not found")
}

// SearchProduct searches for a product by ID or Name
func SearchProduct(query interface{}) (*Product, error) {
	switch v := query.(type) {
	case int:
		for _, product := range inventory {
			if product.ID == v {
				return &product, nil
			}
		}
	case string:
		for _, product := range inventory {
			if strings.EqualFold(product.Name, v) {
				return &product, nil
			}
		}
	default:
		return nil, errors.New("invalid query type")
	}
	return nil, errors.New("product not found")
}

// DisplayInventory displays all products in the inventory
func DisplayInventory() {
	fmt.Printf("%-5s %-20s %-10s %-10s %-15s %-30s\n", "ID", "Name", "Price", "Stock", "Category", "Description")
	fmt.Println(strings.Repeat("-", 90))
	for _, product := range inventory {
		fmt.Printf("%-5d %-20s %-10.2f %-10d %-15s %-30s\n", product.ID, product.Name, product.Price, product.Stock, product.Category, product.Description)
	}
}

// SortInventory sorts products by price or stock
func SortInventory(by string) error {
	switch by {
	case "price":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Price < inventory[j].Price
		})
	case "stock":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Stock < inventory[j].Stock
		})
	default:
		return errors.New("invalid sort key, use 'price' or 'stock'")
	}
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var choice int
		fmt.Println("\nInventory Management System")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Sort Inventory")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id, stock int
			var name, category, description string
			var price float64
			fmt.Print("Enter Product ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Product Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Product Price: ")
			fmt.Scan(&price)
			fmt.Print("Enter Product Stock: ")
			fmt.Scan(&stock)
			fmt.Print("Enter Product Category: ")
			fmt.Scan(&category)
			description, _ = reader.ReadString('\n')
			fmt.Print("Enter Product Description: ")
			description, _ = reader.ReadString('\n')     // Use reader to take multi-line input
			description = strings.TrimSpace(description) // Remove any trailing newline or space
			if err := AddProduct(id, name, price, stock, category, description); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Product added successfully.")
			}
		case 2:
			var id, newStock int
			fmt.Print("Enter Product ID to update stock: ")
			fmt.Scan(&id)
			fmt.Print("Enter new stock value: ")
			fmt.Scan(&newStock)
			if err := UpdateStock(id, newStock); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Stock updated successfully.")
			}
		case 3:
			var query string
			fmt.Print("Enter Product Name or ID to search: ")
			fmt.Scan(&query)
			if id, err := strconv.Atoi(query); err == nil {
				if product, err := SearchProduct(id); err == nil {
					fmt.Printf("Product Found: %+v\n", *product)
				} else {
					fmt.Println("Error:", err)
				}
			} else {
				if product, err := SearchProduct(query); err == nil {
					fmt.Printf("Product Found: %+v\n", *product)
				} else {
					fmt.Println("Error:", err)
				}
			}
		case 4:
			DisplayInventory()
		case 5:
			var sortBy string
			fmt.Print("Enter sort key (price/stock): ")
			fmt.Scan(&sortBy)
			if err := SortInventory(sortBy); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Inventory sorted successfully.")
				DisplayInventory()
			}
		case 6:
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
