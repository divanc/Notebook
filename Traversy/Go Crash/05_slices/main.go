package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare
	fruitArr := [2]string{"Apple", "Orange"}

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:])
}
