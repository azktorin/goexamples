// Package name.
package main

// Imports.
import (
	"fmt" /*Formating lib.*/
)

// Magic goes here.
func Histogram(arr []int) []string {
	var histo []string
	length := len(arr)

	for j := 0; j < length; j++ {
		var w string

		for i := 0; i < arr[j]; i++ {
			w += "*"
		}

		histo = append(histo, w)
	}

	return histo
}

// Main function.
func main() {
	ints := []int{2, 4, 9, 12, 34, 5, 9, 1, 21, 4}

	histogram := Histogram(ints)

	for v := range histogram {
		fmt.Printf("%s\n", histogram[v])
	}
}
