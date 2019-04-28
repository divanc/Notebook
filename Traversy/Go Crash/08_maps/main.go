package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Regina"] = "reginbegin@gmail.com"
	emails["Yura"] = "has job"

	fmt.Println(emails)

	fmt.Println(len(emails))

	delete(emails, "Yura")
	fmt.Println(emails)

	// Declare map
	mails := map[string]string{"Bob": "Hamilton", "Phil": "Valley"}

	fmt.Println(mails)
}
