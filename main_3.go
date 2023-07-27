package main

import (
	"fmt"
)

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("Value of x is:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of x is:", i)
	// }

	names := []string{"A", "B", "C", "D"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println("Value of x is:", names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("the postion at index %v is %v \n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("the postion at index %v \n", value)
		value = "new string"
	}

	fmt.Println(names)
}
