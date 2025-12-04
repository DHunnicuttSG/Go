package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
	Email string
}

func main() {
	contacts := make(map[string]Contact)

	// Add contacts
	contacts["David"] = Contact{"David", "555-1111", "d@email.com"}
	contacts["Maria"] = Contact{"Maria", "555-2222", "m@email.com"}
	contacts["Joe"] = Contact{"Joe", "555-3333", "j@email.com"}

	// Get a contact
	c, found := contacts["Joe"]
	if found {
		fmt.Println("Found:", c.Name, c.Phone, c.Email)
	}

	// Update
	contacts["Joe"] = Contact{"Joe", "555-9999", "new@email.com"}

	// List all
	fmt.Println("\nAll contacts:")
	for _, info := range contacts {
		fmt.Println(info.Name, "-", info.Phone, "-", info.Email)
	}

	// Delete one
	delete(contacts, "Maria")

	fmt.Println("\nAfter delete:")
	for _, info := range contacts {
		fmt.Println(info.Name)
	}
}
