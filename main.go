package main

import (
	"fmt"
	"github.com/Lemmaritzy/go-notes/pkg/datastructures"
	"github.com/Lemmaritzy/go-notes/pkg/datatypes"
	"github.com/Lemmaritzy/go-notes/pkg/flowcontrol"
	"github.com/Lemmaritzy/go-notes/pkg/functions"
)

func main() {

	//Data types- introduction
	datatypes.Variables()
	datatypes.GetValues()
	datatypes.TypeCasting()

	//Flow control (if | switch | for )
	flowcontrol.Ifelsecondition()
	flowcontrol.SwitchCase()
	flowcontrol.ForLoop()

	//Data structures (array | slice | map | struct ))
	datastructures.Arrays()
	datastructures.Slices()
	datastructures.Maps()
	datastructures.Structures()

	//Functions
	fmt.Println(functions.Recursion(6))
	functions.CountDown(12)
	functions.ContainsAnonymous()
	getInfo := functions.ShowStatus()

	fmt.Println(getInfo()) // anonymous function returned, and in main variable behaves like function

	//Closures
	numbIncrease := functions.IncreaseNums(0, 6)
	fmt.Println(numbIncrease())
	fmt.Println(numbIncrease())
	fmt.Println(numbIncrease())

	numbIncrease2 := functions.IncreaseNums(12, 8)
	fmt.Println(numbIncrease2())
	fmt.Println(numbIncrease2())
	fmt.Println(numbIncrease2())
	fmt.Println(numbIncrease2())
}
