package datastructures

import "fmt"

// Arrays array is a collection of similar types for ex. animals, grades, cars, numbers
func Arrays() {

	var simpleArr = [4]int{1, 24, 53, 23}
	fmt.Println("Basic array declaration of integers ", simpleArr)

	var compiledLanguage [4]string
	compiledLanguage[0] = "Go"
	compiledLanguage[1] = "C"
	compiledLanguage[2] = "C++"
	compiledLanguage[3] = "Haskell"

	//shorthand
	interpretedLanguages := [5]string{"javascript", "php", "ruby", "python"}
	fmt.Println(interpretedLanguages)

	//get length and capacity of an array
	fmt.Println(len(interpretedLanguages), cap(interpretedLanguages))

	//multi dimensional arrays

	TwoDimensionalArray := [3][2]int{{2, 4}, {1, 3}, {11, 13}}
	fmt.Println("2D array => ", TwoDimensionalArray)

	// if you want to loop through 2d array, you have to use nested for loops. For example;

	for i := 0; i < len(TwoDimensionalArray); i++ {
		for j := 0; j < len(TwoDimensionalArray[i]); j++ {
			fmt.Println(TwoDimensionalArray[i][j])
		}
	}
}

// Slices Slice same behavior with Arrays, but slices don't have fixed size, that means while adding or removing item, slice's capacity will change automatically,
func Slices() {
	mySlice := []int{1, 2, 3, 4, 5} // a slice, because we didn't put size in the []. Free size no limitation
	//also we can create a slice from an array
	fmt.Println(mySlice)

	myArr := [4]int{4, 5, 7, 1}
	mySlice2 := []int{}
	mySlice2 = myArr[0:2]
	fmt.Println(mySlice2)

	//slice functions
	/*
		append => adds new element
		copy => copies element of slice to another slice
		equal=> compares 2 slice
		len => gets the length of any slice
	*/

	negativeNums := []int{-1, -2, -3}
	positiveNums := []int{5, 7, 3}

	//adds more elements
	negativeNums = append(negativeNums, -12, -24)

	//adds positive nums into negative nums
	negativeNums = append(negativeNums, positiveNums...)

	fmt.Println(negativeNums)

}

// Maps  we can store elements with key->value pairs. don't forget, keys are unique
func Maps() {
	countryCodes := make(map[string]string)
	countryCodes["tr"] = "Turkey"
	countryCodes["UK"] = "United Kingdom"
	countryCodes["FR"] = "France"
	countryCodes["US"] = "United States"

	for countryCode, CountryName := range countryCodes {
		fmt.Printf("%s => %s", countryCode, CountryName)
	}

	//if you want to any element from a map, just use delete() function
	//delete-> first parameter is the map, second one is the key that we want to delete
	delete(countryCodes, "FR")

}

// Structures used to store the variables of different datatype for example think about a Person, person holds name,surname string and age int we can store string and int variables together
func Structures() {

	u1 := User{
		"mert", "yılmaz", 22, "555 555 55 55", "mert@example.com",
	}
	var u2 User

	u2 = User{
		firstName: "peter",
		lastName:  "griffin",
		age:       46,
		phone:     "789 985 34 76",
		mail:      "hihohiho@familyguy.com",
	}
	fmt.Println(u1)
	fmt.Println(u2)

	u3 := Person{
		firstName: "Josh",
		lastName:  "Randor",
		weight:    80,
		height:    1.75,
		BMI: func(weight, height float64) float64 {
			return weight / (height * height)
		},
	}

	fmt.Printf("%s's BMI => %v", u3.firstName, u3.BMI(u3.weight, u3.height))

}

//example struct

type User struct {
	firstName, lastName string
	age                 int
	phone, mail         string
}

// struct that includes function

type calcBMI func(w, h float64) float64

type Person struct {
	firstName, lastName string
	weight, height      float64
	BMI                 calcBMI
}
