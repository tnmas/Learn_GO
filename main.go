package main

import (
	"fmt"
)

func userDetails(firstName string, lastName string, age uint, jobType string) {
	fmt.Printf("\nYour Fullname: %v %v \n", firstName, lastName)
	fmt.Printf("Age: %v \n", age)
	fmt.Printf("Your Job: %v \n", jobType)
	fmt.Println("Congrats, you are a member!")
}

func main() {

	const members int = 5
	activeMembers := members
	var firstName string
	var lastName string
	var age uint
	var jobType string
	var membersArray []string

	for i := 0; i < members; i++ {
		fmt.Println("Enter your Firstname: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your Lastname: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your age: ")
		fmt.Scan(&age)
		fmt.Println("Enter your job: ")
		fmt.Scan(&jobType)

		membersArray = append(membersArray, firstName+" "+lastName)
		activeMembers -= 1

		userDetails(firstName, lastName, age, jobType)
		fmt.Printf("\nNumber of members: %v \n", activeMembers)
	}

	fmt.Println("Active members:", membersArray)
}
