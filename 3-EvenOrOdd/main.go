// In the main function, create a slice of ints from 0 through 10
// Iterate through the slice.  For each element, check to see if the number is even or odd.
// If the value is even, print out "even".  If it is odd, print out "odd"

package main

import "fmt"

func main() {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range values {
		if value%2 == 0 {
			fmt.Println("The number is even", value)

		} else {
			fmt.Println("The number is odd", value)
		}
	}
}
