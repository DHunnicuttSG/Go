package main

import (
	"encoding/json"
	"fmt"
)

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

type JContact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c Contact) fullName() string {
	return c.FirstName + " " + c.LastName
}

func (c *Contact) updatePhone(phone string) {
	c.Phone = phone
}

func main() {

	c := Contact{
		ID:        1,
		FirstName: "David",
		LastName:  "Hunnicutt",
		Phone:     "812-555-1234",
		Email:     "DH@mthree.com",
	}

	fmt.Println(c.fullName())

	fmt.Println("Before Update:", c.Phone)
	// update phone number
	c.updatePhone("555-1111")
	fmt.Println("After Update:", c.Phone)

	jc := JContact{2, "David H", "dh@mthree.com"}

	jsonData, _ := json.Marshal(jc)
	fmt.Println(string(jsonData))

	var c2 JContact
	json.Unmarshal([]byte(jsonData), &c2)
	fmt.Println(c2.Name)

}
