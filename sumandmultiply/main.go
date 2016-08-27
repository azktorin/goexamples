// Package name.
package main

// Imports.
import (
	"fmt" /*Formating lib.*/
)

// Sum function. Return sum of given array. (slice).
func Sum(arr []int) int {
	var s int
	for v := range arr {
		s += arr[v]
	}

	return s
}

// Multiply function. Return multiply of given array. (slice).
func Multiply(arr []int) int {
	var p int = 1
	for v := range arr {
		p *= arr[v]
	}

	return p
}

// Main function.
func main() {
	arr := []int{34, 1, 6, 2}

	s := Sum(arr)
	m := Multiply(arr)

	fmt.Printf("Sum = %d\nMultiply = %d\n", s, m)
}
