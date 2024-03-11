// Write a program that accepts a sentence and calculate the number of letters and digits.

// Suppose the following input is supplied to the program:

// hello world! 123

// Then, the output should be:

// LETTERS 10
// DIGITS 3

package ex10

import "unicode"

func countLettersAndDigits(input string) (int, int) {
	var letters, digits int
	for _, char := range input {
		if unicode.IsLetter(char) {
			letters++
		} else if unicode.IsDigit(char) {
			digits++
		}
	}
	return letters, digits
}