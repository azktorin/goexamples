// Package name.
package main

// Imports.
import (
	"fmt" /*Formating lib.*/
)

// Function that return reversed strings.
func ReverseIt(str string) string {
	var rstr string
	nchars := len(str)

	for i := nchars - 1; i >= 0; i-- {
		rstr += string(str[i])
	}

	return rstr
}

// Main function.
func main() {
	str := "Hi there, there is a new string!"

	reversed := ReverseIt(str)

	fmt.Printf("{%s} = {%s}\n", str, reversed)
}
