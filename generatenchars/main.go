// Package name.
package main

// Imports.
import (
	"fmt" /*Formating lib.*/
)

// Magic here. Function that return a string of "n" chars "c".
func GenerateNChars(n int, c string) string {
	var s string

	for i := 0; i < n; i++ {
		s += c
	}

	return s
}

// Main function.
func main() {
	result := GenerateNChars(9, "z")

	fmt.Printf("Generated chars - %s\n", result)
}
