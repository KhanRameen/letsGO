package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
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

	//Array: make your array capaciy
	var intSlice2 []int32 = make([]int32, 8, 15) //type,length,capacity
	intSlice2 = append(intSlice2, intSlice...)
	fmt.Println(intSlice2)

	//maps

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
