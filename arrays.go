package main

import "fmt"

func main() {
	//All right then so lets move to two different types arrays and slices. We'll start with Arrays then go on to Slices.

	//Arrays
	//How do we create an array, how do we declare an array variable? Well, same way as we would any other variable

	//Let's declare var and give it a name of *carNames* and if we want to explicitly type this we can do it using square brackets.

	// *[]*

	//First to signify. This is going to be an array and inside the square brackets we see how many items are going to be in the array.
	//So if I want three items inside it, it's going to be three.
	// [*3*]

	//Then we say what type of data is going to be inside the array? This now says the carNames will be a number array of three items
	// [3]*int*

	//So let's give this a value. Before defining a value we also have to kind of put the typing on the right-hand side of the expression as well.
	//So on the right, before we actually create the array.
	// var carNames [3]int = [3]int
	// This might seem a bit weird but this is how we do it on Go
	// Then the array itself is curly braces not square brackets'
	// var carNames [3]int = [3]int{1, 3, 4}
	// There's now three of these inside and array. 1,3,4



	var carNames [3]int = [3]int{1, 3, 4}

	fmt.Println(carNames)

}
