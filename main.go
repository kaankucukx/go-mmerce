package main

import "fmt"

func main() {
	name := "Kağan"
	age := 99

	// Printinh things
	fmt.Print("hello, ")
	fmt.Print("boom! \n")
	fmt.Print("new line! \n")

	println("höllööö")
	println("höllööö 2")

	// Formated Strings
	// See how the console output order is :)

	println("formatted sTrings")
	fmt.Printf("age: %v, name: %v", age, name)

}
