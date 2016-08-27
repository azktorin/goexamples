// Package name.
package main

// Imports.
import (
	"fmt"
)

// IsPalindrome function.
func IsPalindrom(word string) bool {
	var rword string    // Make a variable 'rword' - reversed word, where we'll save reversed word.
	nchars := len(word) // Get the length of given word.

	//The magic below. Novice example so we can't using runes.
	for i := nchars - 1; i >= 0; i-- {
		rword += string(word[i])
	}

	if word == rword {
		return true
	}

	return false
}

// Main function.
func main() {
	word := "racecar" /* var word string = "racecar" */

	pal := IsPalindrom(word)

	fmt.Printf("Is '%s' a palindrome? - %t\n", word, pal) // Output.
}
