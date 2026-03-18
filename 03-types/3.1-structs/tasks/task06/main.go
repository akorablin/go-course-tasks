package main

import "fmt"

type ContactInfo struct {
	Phone string
	Email string
}

type Address struct {
	City   string
	Street string
}

type Client struct {
	ID      string
	Address Address
	ContactInfo
}

func main() {
	c := Client{
		ID: "1",
		Address: Address{
			City:   "Moscow",
			Street: "Berzarina",
		},
		ContactInfo: ContactInfo{
			Phone: "111",
			Email: "222",
		},
	}
	fmt.Printf("ID: %s, City: %s, Email: %s\n", c.ID, c.Address.City, c.Email)
}
