package main

import (
	"SampleApp/Utils"
	"fmt"
)

func main() {
	fmt.Println("Hello World, My Name is Go Language!")

	// declare variables
	var myName string = "myself"
	var herName = "someOne"
	var hisName string

	fmt.Println(myName)
	fmt.Println(herName)
	fmt.Println(hisName)

	hisName = "himself"

	fmt.Println(hisName)

	var myAge = 23
	var herAge = 44
	hisAge := 65

	fmt.Println(myAge, herAge, hisAge)

	var nums = [3]int{20, 21, 22}

	fmt.Println(nums, len(nums))

	var iter int= 0
	for iter < 5 {
		fmt.Println("value x: ", iter)
		iter++
	}

	ages := []int {myAge, herAge, hisAge}

	for index, value := range  ages {
		fmt.Println(index, value)

		if value < 30 {
			fmt.Println("less than 30")
		} else if value < 50 {
			fmt.Println("less than 50 but more than 30")
		} else {
			fmt.Println("greater than 50")
		}
	}

	greet := Utils.GetGreeting("fatih")
	fmt.Println(greet)
	numOne, numTwo := Utils.GetTwoNextNumbers(myAge)
	fmt.Println(numOne, numTwo)

	// Non Pointer Values
	code := 3
	Utils.UpdateIntValue(code)
	fmt.Println(code)

	// Pointer Wrapper Values
	fruitDict := map[string]int{
		"orange": 2,
		"mango": 4,
	}
	Utils.UpdateMapValue(fruitDict, "orange", 8)
	fmt.Println(fruitDict["orange"])
}

