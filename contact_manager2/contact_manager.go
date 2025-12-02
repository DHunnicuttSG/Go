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

// ----------------- MAIN MENU -----------------

func main() {

	// Add a contact
	addContact(Contact{1, "Ramiz", "Abdulla", "Initech", "123-456-7890", "RA@email.com"})                  // Office Space
	addContact(Contact{2, "Theon", "Beckford", "Stark Industries", "333-888-2222", "TB@email.com"})        // Iron man
	addContact(Contact{3, "Alanna", "Carton", "Wayne Enterprises", "333-888-2222", "AC@email.com"})        // Batman
	addContact(Contact{4, "Liz", "Coles", "Monsters, Inc.", "333-888-2222", "LC@email.com"})               // Monster's Inc.
	addContact(Contact{5, "Joe", "Haslam", "Genco Olive Oil Company", "333-888-2222", "JH@email.com"})     // Godfather
	addContact(Contact{6, "Azizfatima", "Hussain", "Cyberdyne Systems", "333-888-2222", "AH@email.com"})   // Terminator
	addContact(Contact{7, "Kiran", "Mamidala", "Apex Dynamix", "333-888-2222", "KM@email.com"})            // Company from video games
	addContact(Contact{8, "Nikhitha", "Naik", "Tyrell Corporation", "333-888-2222", "NN@email.com"})       // Blade Runner
	addContact(Contact{9, "Nicole", "Samuels", "Dunder Mifflin", "333-888-2222", "NS@email.com"})          // The Office
	addContact(Contact{10, "Uzaer", "Shahid", "Hooli", "333-888-2222", "US@email.com"})                    // Silicon Valley
	addContact(Contact{11, "Ellis", "Stonehouse", "Umbrella Corporation", "333-888-2222", "ES@email.com"}) // Resident Evil games
	addContact(Contact{12, "Charlie", "Wilson", "Sterling Cooper", "333-888-2222", "CW@email.com"})        // Mad men series

	for {
		fmt.Println("\n===== CONTACT MANAGER =====")
		fmt.Println("1. View all contacts")
		fmt.Println("2. Add a contact")
		fmt.Println("3. Update a contact")
		fmt.Println("4. Delete a contact")
		fmt.Println("5. Find contact by ID")
		fmt.Println("6. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			printContacts(getAllContacts())

		case 2:
			addContactMenu()

		case 3:
			updateContactMenu()

		case 4:
			deleteContactMenu()

		case 5:
			findContactMenu()

		case 6:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

// ----------------- FILE FUNCTIONS -----------------

func loadContacts() []Contact {
	var contacts []Contact

	data, err := os.ReadFile(fileName)
	if err != nil {
		return contacts
	}

	json.Unmarshal(data, &contacts)
	return contacts
}

func saveContacts(contacts []Contact) {
	data, _ := json.MarshalIndent(contacts, "", "  ")
	os.WriteFile(fileName, data, 0644)
}

// ----------------- CRUD FUNCTIONS -----------------

func getAllContacts() []Contact {
	return loadContacts()
}

func addContact(newContact Contact) {
	contacts := loadContacts()
	contacts = append(contacts, newContact)
	saveContacts(contacts)
}

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

func getNextID() int {
	contacts := loadContacts()

	maxID := 0
	for _, c := range contacts {
		if c.ID > maxID {
			maxID = c.ID
		}
	}

	return maxID + 1
}

// ----------------- MENU FUNCTIONS -----------------

func addContactMenu() {
	var c Contact

	c.ID = getNextID()

	fmt.Print("First name: ")
	fmt.Scanln(&c.FirstName)

	fmt.Print("Last name: ")
	fmt.Scanln(&c.LastName)

	fmt.Print("Company: ")
	fmt.Scanln(&c.Company)

	fmt.Print("Phone: ")
	fmt.Scanln(&c.Phone)

	fmt.Print("Email: ")
	fmt.Scanln(&c.Email)

	addContact(c)
	fmt.Println("✅ Contact added")
}

func updateContactMenu() {
	var id int
	var phone string

	fmt.Print("Enter Contact ID to update: ")
	fmt.Scanln(&id)

	fmt.Print("Enter new phone: ")
	fmt.Scanln(&phone)

	updateContact(id, phone)
	fmt.Println("✅ Contact updated")
}

func deleteContactMenu() {
	var id int

	fmt.Print("Enter Contact ID to delete: ")
	fmt.Scanln(&id)

	deleteContact(id)
	fmt.Println("✅ Contact deleted")
}

func findContactMenu() {
	var id int
	fmt.Print("Enter Contact ID: ")
	fmt.Scanln(&id)

	contacts := loadContacts()
	for _, c := range contacts {
		if c.ID == id {
			fmt.Println("\n--- Contact Found ---")
			fmt.Printf("ID: %d\nName: %s %s\nCompany: %s\nPhone: %s\nEmail: %s\n",
				c.ID, c.FirstName, c.LastName, c.Company, c.Phone, c.Email)
			return
		}
	}

	fmt.Println("❌ Contact not found")
}

// ----------------- DISPLAY -----------------

func printContacts(contacts []Contact) {
	if len(contacts) == 0 {
		fmt.Println("No contacts found.")
		return
	}

	fmt.Println("\n---- CONTACT LIST ----")
	for _, c := range contacts {
		fmt.Printf("ID: %d | %s %s | %s | %s | %s\n",
			c.ID, c.FirstName, c.LastName,
			c.Company, c.Phone, c.Email)
	}
}
