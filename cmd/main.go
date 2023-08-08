package main

import (
	"fmt"
	"gogogadget/pkg/tools"
)

func mainJson() {
	jsonStr := `{
		"name": "John Doe",
		"age": 30,
		"country": "USA",
		"contacts": {
			"email": "john@example.com",
			"phone": "123-456-7890"
		},"orders": [
			{
				"order_id": "12345",
				"amount": 100.50
			},
			{
				"order_id": "54321",
				"amount": 75.20
			}
		]
	}`
	tabSize := 4

	formattedJSON, err := tools.IndentJSON(jsonStr, tabSize)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(formattedJSON)
}

func mainPassword() {
	options := tools.PasswordOptions{
		Length:     16,
		NumDigits:  4,
		NumSymbols: 2,
		// AllowUppercase: true,
		// AllowLowercase: true,
		// AllowDigits:    true,
		// AllowSymbols:   true,
	}

	password, err := tools.GeneratePassword(options)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Generated password:\n" + password)
}

func main() {
	// mainJson()
	mainPassword()
	// StartCli()
}
