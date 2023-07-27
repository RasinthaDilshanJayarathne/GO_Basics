package main

import (
	"fmt"
	"sort"
)

func main_3() {
	//greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "th"))
	// fmt.Println(strings.Split(greeting, " "))

	//the original value is unchanged
	// fmt.Println("Original string value =", greeting)

	ages := []int{45, 20, 30, 50, 60}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 60)
	fmt.Println(index)
}
