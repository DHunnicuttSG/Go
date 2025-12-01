package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const fileName = "contacts.json"

type Contact struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func main() {

	// Add a contact
	addContact(Contact{1, "David", "Hunnicutt", "ACME Corp", "123-456-7890", "david@email.com"})
	addContact(Contact{2, "Sarah", "Smith", "Tech Ltd", "333-888-2222", "sarah@email.com"})
	addContact(Contact{2, "Sarah", "Smith", "Tech Ltd", "333-888-2222", "sarah@email.com"})
	addContact(Contact{2, "Sarah", "Smith", "Tech Ltd", "333-888-2222", "sarah@email.com"})
	addContact(Contact{2, "Sarah", "Smith", "Tech Ltd", "333-888-2222", "sarah@email.com"})
	addContact(Contact{2, "Sarah", "Smith", "Tech Ltd", "333-888-2222", "sarah@email.com"})
	addContact(Contact{2, "Sarah", "Smith", "Tech Ltd", "333-888-2222", "sarah@email.com"})
	addContact(Contact{2, "Sarah", "Smith", "Tech Ltd", "333-888-2222", "sarah@email.com"})
	addContact(Contact{2, "Sarah", "Smith", "Tech Ltd", "333-888-2222", "sarah@email.com"})
	addContact(Contact{2, "Sarah", "Smith", "Tech Ltd", "333-888-2222", "sarah@email.com"})
	addContact(Contact{2, "Sarah", "Smith", "Tech Ltd", "333-888-2222", "sarah@email.com"})
	addContact(Contact{2, "Sarah", "Smith", "Tech Ltd", "333-888-2222", "sarah@email.com"})

	// List all contacts
	fmt.Println("\n--- All Contacts ---")
	contacts := getAllContacts()
	printContacts(contacts)

	// Update a contact
	fmt.Println("\n--- Updating Contact 1 ---")
	updateContact(1, "999-999-9999")
	printContacts(getAllContacts())

	// Delete a contact
	fmt.Println("\n--- Deleting Contact 2 ---")
	deleteContact(2)
	printContacts(getAllContacts())
}

//////////////////////////////
// FILE FUNCTIONS
//////////////////////////////

func loadContacts() []Contact {
	var contacts []Contact

	data, err := os.ReadFile(fileName)
	if err != nil {
		return contacts // return empty if file doesn't exist
	}

	json.Unmarshal(data, &contacts)
	return contacts
}

func saveContacts(contacts []Contact) {
	data, _ := json.MarshalIndent(contacts, "", "  ")
	os.WriteFile(fileName, data, 0644)
}

//////////////////////////////
// CRUD FUNCTIONS
//////////////////////////////

// CREATE
func addContact(newContact Contact) {
	contacts := loadContacts()
	contacts = append(contacts, newContact)
	saveContacts(contacts)
}

// READ
func getAllContacts() []Contact {
	return loadContacts()
}

// UPDATE
func updateContact(id int, newPhone string) {
	contacts := loadContacts()

	for i := 0; i < len(contacts); i++ {
		if contacts[i].ID == id {
			contacts[i].Phone = newPhone
			break
		}
	}

	saveContacts(contacts)
}

// DELETE
func deleteContact(id int) {
	contacts := loadContacts()
	var updated []Contact

	for _, c := range contacts {
		if c.ID != id {
			updated = append(updated, c)
		}
	}

	saveContacts(updated)
}

//////////////////////////////
// OUTPUT
//////////////////////////////

func printContacts(contacts []Contact) {
	for _, c := range contacts {
		fmt.Printf("ID: %d | %s %s | %s | %s | %s\n",
			c.ID, c.FirstName, c.LastName,
			c.Company, c.Phone, c.Email)
	}
}
