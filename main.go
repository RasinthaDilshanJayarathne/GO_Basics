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

	age := 25
	name := "Rasintha"

	//print
	fmt.Print("hello, ")
	fmt.Print("world! \n ")
	fmt.Print("new line \n")

	//println
	fmt.Println("hello rasintha!")
	fmt.Println("goodbyee rasintha!")
	fmt.Println("my age is", age, "and my name is", name)

	//printf (Formatting Strings)
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("your scored %f points! \n", 225.55)
	fmt.Printf("your scored %0.1f points! \n", 225.55)

	//Sprintf (Save Formatting Strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string..", str)
}
