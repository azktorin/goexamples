// Package name.
package main

// Imports.
import (
	"fmt" /*Formating lib.*/
)

// Magic function. Very basic and simple.
func Overlapping(l1, l2 []string) bool {
	for v := range l1 {

		for k := range l2 {

			if l1[v] == l2[k] {
				return true
			}
		}
	}

	return false
}

// Main function.
func main() {
	list1 := []string{"Caesar", "August", "Ceceron", "Iulius"}
	list2 := []string{"Andrew", "Ann", "Ben", "Jack", "Iulius"}

	result := Overlapping(list1, list2)

	fmt.Printf("Are there members in common? - %t\n", result)
}
