package main

import (
	"fmt"
)

func userDetails(firstName string, lastName string, age uint, jobType string) {
	fmt.Printf("\nYour Fullname: %v %v \n", firstName, lastName)
	fmt.Printf("Age: %v \n", age)
	fmt.Printf("Your Job: %v \n", jobType)

}

func main() {
	var firstName string
	var lastName string
	var age uint
	var jobType string

	fmt.Println("Enter your Firstname: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Lastname: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your age: ")
	fmt.Scan(&age)
	fmt.Println("Enter your job: ")
	fmt.Scan(&jobType)

	userDetails(firstName, lastName, age, jobType)
}
