package flowcontrol

import "fmt"

func Ifelsecondition() {
	//if else condition

	var isActive bool = true

	// or shorty if isActive {code}
	if isActive == true {

		fmt.Println("active user")
	} else {
		fmt.Println("inactive user")
	}

	//if else statement with (else if) -- ladder if else statement
	var name string = "mert"

	if name == "mert" {
		fmt.Println("you are allowed to system")
	} else if name == "dert" {
		fmt.Println("you are allowed too but dont have 100% access")
	} else {
		fmt.Println("you are not allowed to system")
	}

}

// SwitchCase you can execute code block among many alternatives
func SwitchCase() {
	var dayIndex int = 3

	switch dayIndex {
	case 0:
		fmt.Println("Monday")
	case 1:
		fmt.Println("tuesday")
	case 2:
		fmt.Println("wednesday")
	case 3:
		fmt.Println("thursday")
	case 4:
		fmt.Println("friday")
	case 5:
		fmt.Println("saturday")
	case 6:
		fmt.Println("sunday")
	default:
		fmt.Println("what the hell? there is no day name with this value")
	}

	//we can use fallthrough if we want to still execute after matching value
	//just type fallthrough
	//also you can give multiple values in case like case 1,2,3,4:
}

func ForLoop() {
	var sum int
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Printf("sum from for loop => %v\n", sum)

	//there is one more for loop range. we can look through an array or slice or map, like for each
	//Range

	users := [4]string{"team", "city", "note", "excalibur"}
	for index, item := range users {
		fmt.Printf("index => %s - value => %v", item, index)
		//if you don't want to index, just replace 'index' with blank identifier '_'
	}

}

/*
Break and continue
brake=> terminates the loop when it's encountered
continue => skips the current iteration, still works till end

*/
