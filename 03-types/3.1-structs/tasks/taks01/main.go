package main

import "fmt"

type Profile struct {
	Name     string
	Age      int
	IsActive bool
}

func main() {
	p := Profile{
		Name:     "Alexander",
		Age:      34,
		IsActive: true,
	}

	fmt.Printf("%+v\n", p)
}
