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

	pointers()
}

func pointers() {
	//Pointers: they have their own memory address and on their memory address they store memory address of others. eg: [0x1b0c]@0x1b00

	var p *int32 //stores memory addressd of 32 bits , right now is nill (doesnt point anything). Throws a run time error when called due to nill
	var i int32  //stores default value 0 of 32 bits

	var newPointer *int32 = new(int32) //bascially it now stores a memory location an not nill so whatever the value will be on that memory location it will point to it. rightnow it is 0 (default) ;
	//newPointer: [0x1b0c]@0x1b00
	//new(int32): [0]@0x1b0c    ^

	var pointerForI *int32 = &i
	p = &i

	fmt.Printf("The value p points to is %v\n", *p)
	fmt.Printf("The value newPointer points to is %v\n", *newPointer)
	fmt.Printf("The value PointerForI is %v\n", pointerForI)

	*p = 15
	fmt.Printf("The value p after i derefrenced it and chaged the value on the memory location it store (var i): %v\n", *p)
	fmt.Printf("The value pointerForI after i derefrenced p and chaged the value on the memory location it store (also var i): %v\n", *pointerForI)

	//slices : slices themselves are pointers to an underlying array. yeah. digest that or whatever
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println("SLice: ", slice)
	fmt.Println("SLice Copy: ", sliceCopy)
	//they both point to same array location so changing one changes both

	thing1 := [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("Memory Location of Thing1 is %p\n", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("Square Result is %v\n", result)
	fmt.Printf("The Value for Thing1 now is %v\n", thing1)

}

// pointers in function: usually you pass the value to the arg variable in the funciton and it returns a value in a new location. with points you can pass the refernce and so you can modify the actual value on the memory location passed without wasting a new space. (this obviously modifies the original passed value so you gotta know you use case)

func square(thing2 *[5]float64) [5]float64 { //takes pointer to an array not an array itself.
	fmt.Printf("The memory location for thing2 array is: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2 //the value on thing2
}
