package main

import "fmt"

func test(name string)  {
	fmt.Println("Hello", name)
}

func main() {
	yourName := "John Smith"
	test(yourName)
}