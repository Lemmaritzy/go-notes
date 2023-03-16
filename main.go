package main

import (
	"github.com/Lemmaritzy/go-notes/pkg/datastructures"
	"github.com/Lemmaritzy/go-notes/pkg/datatypes"
	"github.com/Lemmaritzy/go-notes/pkg/flowcontrol"
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
}
