package main

import "fmt"

type Address struct {
	City   string
	Street string
}

type Employee struct {
	Name    string
	Address Address
}

func main() {
	u := Employee{
		Name: "Alexander",
		Address: Address{
			City:   "Moscow",
			Street: "Berzarina",
		},
	}
	fmt.Printf("%s: %s, %s\n", u.Name, u.Address.City, u.Address.Street)
}
