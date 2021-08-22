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
