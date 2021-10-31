package main

import "fmt"

func main() {

	// Count int variable assigned the num 42
	var count = int(42)

	// ptr is assgned the address where the count variable is located
	ptr := &count

	// Prints the count value that is pointed to by the *ptr variable
	fmt.Println(*ptr)

	// Assigns a new value to the address where the count variable is located
	*ptr = 100

	// Prints the new value
	fmt.Println(count)
}
