package main

import (
	"fmt"
)

// structs
type nameOfType struct {
	name         string
	val          uint8
	officialInfo official
	int
}

type official struct {
	isOfficial bool
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

// methods: functions of structs
func (e gasEngine) milesLeft() uint8 { // (varName struct) nameOfMethod() returntype
	return e.gallons * e.mpg
}

// interface: methods general to any struct
type engineOfAnyType interface {
	milesLeft() uint8
}

func canMakeIt(e engineOfAnyType, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there")
	} else {
		fmt.Println("Oops you can't, you need more fuel")
	}
}

func notSoBasics() {
	var myStruct nameOfType                  //default values
	fmt.Println(myStruct.name, myStruct.val) // returns "", 0
	//or

	var structWVal nameOfType = nameOfType{"Fin", 10, official{true}, 14}
	fmt.Println(structWVal.name, structWVal.val, structWVal.officialInfo, structWVal.int)

	//anonymous struct: no name type defined , only brackets. Not Reusable.
	var weirdStruct = struct {
		name string
		age  uint8
	}{"Adam", 90}
	fmt.Println(weirdStruct.name, weirdStruct.age)

	var myEngine gasEngine = gasEngine{25, 15}
	canMakeIt(myEngine, 50)
}
