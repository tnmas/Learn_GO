package main

import "fmt"

//var yourName string

func test(name string)  {
	fmt.Println("Hello", name)
}

func main() {
	yourName := "John Smith"
	test(yourName)
}