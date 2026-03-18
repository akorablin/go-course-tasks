package main

import "fmt"

type Package struct {
	ID     string
	Weight int
}

type Destination struct {
	City string
	Zip  string
}

type Shipment struct {
	Package     Package
	Destination Destination
}

func main() {
	s := Shipment{
		Package: Package{
			ID:     "Test",
			Weight: 100,
		},
		Destination: Destination{
			City: "Moscow",
			Zip:  "3420845",
		},
	}
	fmt.Printf("ID: %s, City: %s\n", s.Package.ID, s.Destination.City)
}
