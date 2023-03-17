package functions

import (
	"fmt"
)

//We're using functions to divide our code to necessary and small blocks of code. So our code becomes more clean and readable. We are organizing our code

func Greetings() {
	fmt.Println("Hello there!")
}

func AreaOfSquare(x int) int {
	area := x * x
	return area
}

// if you want to create public function, uppercase first letter. so if want to create private function, just type first character lower case

func PublicFunction() {
	fmt.Println("You can call me where ever you want to ")
}

func privateFunction() {
	fmt.Println("I'm only available in function package!")
}

// Recursion functions meaning recursive functions calls itself
func Recursion(x int) int {

	if x > 0 {
		return x * Recursion(x-1)
	} else {
		return 1
	}
}

func CountDown(x int) {

	if x > 0 {
		fmt.Printf("Counting down from %v\n", x)

		CountDown(x - 1)
	} else {
		fmt.Println("end")
	}
}

// ContainsAnonymous functions = we are assigning our unnamed function to variable
func ContainsAnonymous() {
	fmt.Println("Example for anonymous function")

	call := func(x, y int) int {
		return x * y
	}
	fmt.Println("Created anonymous function and assigned to call(variable). ", call(5, 6))

}

//Also you can return an anonymous function in another function or another package if parent of anonymous function is public
// You can use anonymous functions as argument

func ShowStatus() func() string {
	var num = 1
	var text string
	switch num {
	case 1:
		text = "Correct!"
	case 2:
		text = "False!"
	case 3:
		text = "False"
	case 4, 5, 6:
		text = "False"
	default:
	}

	return func() string {
		text = text + " Ow yeaah"
		return text
	}
}

// Closures is a nested functions that allow us to access variables of the outer function

func IncreaseNums(index, beh int) func() int {
	num := index

	return func() int {
		num = num + beh
		return num
	}

}
