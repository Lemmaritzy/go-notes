package datatypes

import "fmt"

//declaring variables out of scope

var var4 int = 100
var var5 int

// const are the fixed values, you cant change them
const userType = 2

// you can not use shorthand declaring out of scope
// var6:=12 == ERROR

func Variables() {

	//declaring variables

	var var1 string = "var 1"
	var var2 = 10 //automatically detects the type
	var3 := 12.25 //shorthand declaring

	//creating multiple variables
	var name, surname string
	name = "mert"
	surname = "yılmaz"
	mail, number := "mert@examplemail.com", "1111"

	fmt.Printf("in scope variables;\n first => %s\n second => %v\n third => %v\n", var1, var2, var3)
	fmt.Println("-----------------------------------------------------")

	fmt.Printf("out of scope variables; \n var4 => %v\n var5 => %v\n", var4, var5)
	fmt.Println("-----------------------------------------------------")

	fmt.Printf("User informations; \n name=> %s\n surname=> %s\n mail =>%s\n number %s\n user type => %v\n", name, surname, mail, number, userType)
	fmt.Println("-----------------------------------------------------")

	//Data types
	/*

	   Integer types
	   int-uint(8-16-32-64)

	   Float types
	   float32,float64

	   String type
	   string (array of chars)

	   Boolean type
	   bool (true or false)

	*/

}

// GetValues for taking values via cmd
func GetValues() {
	var name string
	fmt.Println("What is your name? ")
	fmt.Scanln(&name)

	fmt.Printf("Hi %s", name)
}

// TypeCasting example for converting a type to another type
func TypeCasting() {

	var float_num float32 = 12.1
	var int_num int = 24

	//we converted int num to float for addition
	sum := float_num + (float32(int_num))
	fmt.Printf("float 12.1 plus int 24 => %v and type %g", sum, sum)
}
