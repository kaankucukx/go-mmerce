package main

import "fmt"

func main() {
	//Arrays

	var things [3]int = [3]int{1, 3, 4}

	// Don't allow us to create four length array
	//var things [3]int = [3]int{1, 3, 4}

	// Don't let us declare string
	//var things [3]int = [3]int{1, 3, "string"}

	//Displaying length of an array
	number := [3]int{1, 2, 3}
	//fmt.Println(things, number, len(number))

	// Computed properties for defining arrays with shorthand variable.
	//"d" is not allowed since number len is 3
	names := [len(number)]string{"a", "b", "c"}

	fmt.Println(things, names)
}
