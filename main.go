package main

import "fmt"

func main() {
	name := "Kağan"
	age := 99

	// Printf formats according to a format specifier and writes to standard output.
	// It returns the number of bytes written and any write error encountered.

	// Printinh things
	fmt.Print("hello, ")
	fmt.Print("boom! \n")
	fmt.Print("new line! \n")

	println("höllööö")
	println("höllööö 2")

	// Formated Strings
	// See how the console output order is :)

	println("formatted sTrings")
	fmt.Printf("age: %d, name: %q \n", age, name)

	//Logging type of the var
	fmt.Printf("type of age is %T \n", age)

	//Param logging
	fmt.Printf("log the given %f variable \n", 999.99)
	fmt.Printf("log the given %0.1f variable \n", 999.99)
	fmt.Printf("log the given %0.11f variable \n", 999.99)

	//fmt.Sprintf() Save formatted strings
	var test = fmt.Sprintf("bla bla %v", name)
	println("saved string is", test)

	//https://pkg.go.dev/fmt
}
