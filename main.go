package main

import (
	"fmt"
)

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


	var carNames [3]int = [3]int{1, 3, 4}


	// Don't allow us to create four length array
	//var carNames [3]int = [3]int{1, 3, 4}

	// Don't let us declare string
	//var carNames [3]int = [3]int{1, 3, "string"}

	//Displaying length of an array
	number := [3]int{1, 2, 3}
	//fmt.Println(carNames, number, len(number))

	// Computed properties for defining arrays with shorthand variable.
	//"d" is not allowed since number len is 3
	names := [len(number)]string{"a", "b", "c"}

	fmt.Println(carNames, names)

	// Slices (use Arrays)
	//You can change the value of a property of both Array and Slices with targeting index [1] and you can return new Slice by using append fn. Append does not work on Arrays

	var goals = []int{3, 5, 6}

	fmt.Println("goals:", goals)
	goals[2] = 10

	newSliceByAppendingValue := append(goals, 90)

	fmt.Println("goals:", goals, "newSliceByAppendingValue:", newSliceByAppendingValue, "newSliceByAppendingValue length:", len(newSliceByAppendingValue))

	//Slice Ranges

	fmt.Println("Slice Ranges")

	firstRange := goals[0:2]
	secondRange := goals[1:]
	thirdRange := goals[:2]

	fmt.Println(firstRange, secondRange, thirdRange)
}
