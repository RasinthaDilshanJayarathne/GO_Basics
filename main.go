package main

import "fmt"

func main() {
	//strings
	var nameOne string = "Rasi"
	var nameTwo = "Dilshan"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "123"
	nameThree = "Jaya"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Hello"
	fmt.Println(nameFour)

	//ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 256
	fmt.Println(numOne, numTwo, numThree)

	//float
	var scoreOne float32 = 234.43
	var scoreTwo float64 = 255.334534534534
	scoreThree := 1.5
	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
