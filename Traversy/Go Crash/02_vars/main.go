package main

import "fmt"

func main() {

	// Using var
	//var name = "Regin"
	name := "Regin"
	var age = 37
	const isCool = true
	size := 1.3

	user, password := "Ivan", "ReginPekin"

	fmt.Println(name, age, isCool, user, password)
	fmt.Printf("%T\n", size)
}
