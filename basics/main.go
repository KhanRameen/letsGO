package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// basics()
	notSoBasics()
}

func basics() {

	fmt.Println("Let's Go!")

	var intNum int16 = 32767
	fmt.Println(intNum)
	fmt.Println(intNum + 1) //-32768 : overflow int16 bit

	var newstring string = `long string 
	with new line`
	fmt.Println(len(newstring)) //read no.of bytes not character :)
	fmt.Println(utf8.RuneCountInString(newstring))

	newvar := "text" //shorthand for var
	var1, var2 := 1, 2
	fmt.Println(newvar, var1, var2)

	print(newstring)

	var result, remainder, err = intDivison(11, 7)
	if err != nil { //doable with switches as well
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Println("Division Result is ", result)
	} else {
		fmt.Println("Division Result is ", result, " and the remainder is ", remainder)
	}

	//Arrys: fixed size
	var intArray [3]int32 //index space of 0 1 2, 4 bytes each
	intArray[1] = 123     //assign
	fmt.Println(intArray[1:3])
	fmt.Println("Memory location", &intArray[1]) //contiguous

	var anotherArray [2]int16 = [2]int16{4, 5} //immediate initialization
	fmt.Println(anotherArray)

	andAnotherArray := [2]string{"hehe", "huhu"}
	print(andAnotherArray[0])

	//Arrays:Slices for more convinience
	var intSlice []int32 = []int32{2, 4, 5, 7, 6}
	fmt.Printf("Before append, length: %v with capacity: %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 10)                                                             //adds at the end
	fmt.Printf("\nAfter append, length: %v with capacity: %v \n", len(intSlice), cap(intSlice)) //adds extra capacity of 2

	//Array: make your array capactiy
	var intSlice2 []int32 = make([]int32, 8, 15) //type,length,capacity
	intSlice2 = append(intSlice2, intSlice...)
	fmt.Println(intSlice2)

	//maps: objects
	var newMap map[string]uint8 = make(map[string]uint8) //key of type string with unsigned int of 8 bits
	fmt.Println(newMap)

	var myMap2 = map[string]uint8{"Adam": 45, "Sarah": 22}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["RandomName"]) //you get default value of the value type, in this case 0 :)
	var age, ok = myMap2["Sarah"]     //var ok returns true if the value exist
	fmt.Println("Sarah exist in the map: ", ok)
	if ok {
		fmt.Printf("Sarah's age is %v\n", age)
	}

	delete(myMap2, "Adam") //delete maps value by reference, has no return value

	//Loops
	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	} //iterating over map has no order, its a pure guessing game for which will come first

	for name, age := range myMap2 {
		fmt.Printf("%v is %v year old\n", name, age)
	}

	for i, v := range intSlice {
		fmt.Printf("Value at array index %v is %v\n", i, v)
	}

	//Fancy Shmancy Loops: you get one loop in go. you gotta deal with..
	i := 0
	for i < 5 { //like while loops
		fmt.Println(i)
		i = i + 1
	}
	//or
	for { //also like while
		if i == 0 {
			break
		}
		fmt.Println("while backwards", i)
		i--
	}
	for n := 0; n < 3; n++ {
		fmt.Println("n loop")
	}

	Strings()
}

func Strings() {
	var myString = "Résumé" //utf8,
	var indexed = myString[1]
	fmt.Printf("%v , type: %T\n", indexed, indexed) //return the byte size of 1 index space , é takes 2 bytes so in expands over 2 indexes
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Println("Length of String:", len(myString))

	//use of runes
	var myStringOfRune = []rune("Résumé") //utf8,
	var indexedRune = myStringOfRune[1]
	fmt.Printf("%v , Rune type: %T\n", indexedRune, indexedRune) //return the byte size of 1 index space , é takes 2 bytes so in expands over 2 indexes
	for i, v := range myStringOfRune {
		fmt.Println(i, v)
	}
	fmt.Println("Length of Rune:", len(myStringOfRune))

	var myRune = 'a' //this decalres a rune typed char
	fmt.Printf("\nmyRune = %v", myRune)

	//Strings (from the package)
	var strSlice = []string{"s", "t", "r", "i", "n", "g"}
	var stringBuidler strings.Builder
	for i := range strSlice {
		stringBuidler.WriteString(strSlice[i])
	}
	var catString = stringBuidler.String()
	fmt.Println("\ncatinated:", catString)

}

func print(printValue string) { //because fmt just isnt it
	fmt.Println(printValue)
}

func intDivison(numerator int, denominator int) (int, int, error) {
	var err error //type error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero") //imported class
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder = numerator % denominator
	return result, remainder, err
}
