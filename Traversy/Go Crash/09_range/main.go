package main

import "fmt"

func main() {
	ids := []int{123, 23545, 3213, 5435345, 234}

	for i, id := range ids {
		fmt.Printf("%d — ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("only ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("\nSum = ", sum)

	// Range with map
	mails := map[string]string{"Bob": "Hamilton", "Phil": "Valley"}

	for key, value := range mails {
		fmt.Printf("%s: %s\n", key, value)
	}
}
