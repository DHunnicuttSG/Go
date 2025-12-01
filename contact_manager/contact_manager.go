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
	addContact(Contact{1, "Ramiz", "Abdulla", "Initech", "123-456-7890", "RA@email.com"})
	addContact(Contact{2, "Theon", "Beckford", "Stark Industries", "333-888-2222", "TB@email.com"})
	addContact(Contact{2, "Alanna", "Carton", "Wayne Enterprises", "333-888-2222", "AC@email.com"})
	addContact(Contact{2, "Liz", "Coles", "Monsters, Inc.", "333-888-2222", "LC@email.com"})
	addContact(Contact{2, "Joe", "Haslam", "Genco Olive Oil Company", "333-888-2222", "JH@email.com"})
	addContact(Contact{2, "Azizfatima", "Hussain", "Cyberdyne Systems", "333-888-2222", "AH@email.com"})
	addContact(Contact{2, "Kiran", "Mamidala", "Apex Dynamix", "333-888-2222", "KM@email.com"})
	addContact(Contact{2, "Nikhitha", "Naik", "Tyrell Corporation", "333-888-2222", "NN@email.com"})
	addContact(Contact{2, "Nicole", "Samuels", "Dunder Mifflin", "333-888-2222", "NS@email.com"})
	addContact(Contact{2, "Uzaer", "Shahid", "Hooli", "333-888-2222", "US@email.com"})
	addContact(Contact{2, "Ellis", "Stonehouse", "Umbrella Corporation", "333-888-2222", "ES@email.com"})
	addContact(Contact{2, "Charlie", "Wilson", "Sterling Cooper", "333-888-2222", "CW@email.com"})

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
