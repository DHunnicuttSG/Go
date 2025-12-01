package main

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
	addContact(Contact{1, "Ramiz", "Abdulla", "mthree", "123-456-7890", "RA@mthree.com"})
}
