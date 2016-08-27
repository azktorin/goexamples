// Package name.
package main

// Imports.
import (
	"fmt" /*Formating lib.*/
)

// "RETURN" max of a, b, c.
func MaxOfThree(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else {
		return c
	}
}

// Main function.
func main() {
	max := MaxOfThree(5, 6, 8)

	fmt.Printf("Max of 5, 6, 8 is '%d'\n", max)
}
