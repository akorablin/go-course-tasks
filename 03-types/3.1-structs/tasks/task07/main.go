package main

import "fmt"

type Student struct {
	ID          int
	Login       string
	ContactInfo ContactInfo
}

type Course struct {
	Author string
	Name   string
}

type ContactInfo struct {
	Email string
	Phone string
}

type CourseEnrollment struct {
	EnrollmentID int
	Status       string
	Student      Student
	Course
}

func main() {
	c := CourseEnrollment{
		EnrollmentID: 1,
		Status:       "Start",
		Student: Student{
			ID:    100,
			Login: "test",
			ContactInfo: ContactInfo{
				Email: "Email",
				Phone: "Phone",
			},
		},
		Course: Course{
			Author: "Author",
			Name:   "Name",
		},
	}
	fmt.Printf("Course: %s, Student: %d, Status: %s, EnrollmentID: %d\n", c.Name, c.Student.ID, c.Status, c.EnrollmentID)
}
