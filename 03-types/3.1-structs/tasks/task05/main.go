package main

import "fmt"

type Audit struct {
	CreatedAt string
	UpdatedAt string
}

type Article struct {
	Title string
	Audit
}

func main() {
	a := Article{
		Title: "Test",
		Audit: Audit{
			CreatedAt: "1",
			UpdatedAt: "2",
		},
	}
	fmt.Printf("Title: %s, CreatedAt: %s, UpdatedAt: %s\n", a.Title, a.CreatedAt, a.UpdatedAt)
}
