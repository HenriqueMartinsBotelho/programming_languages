// Write a program that calculates and prints the value according to the given formula:

// Q = Square root of [(2 * C * D)/H]

// Following are the fixed values of C and H:

//     C is 50. H is 30.

//     D is the variable whose values should be input to your program in a comma-separated sequence.

package ex6

import "math"

const c = 50
const h = 30

func Ex006(d []int) []int {
	var s = make([]int, len(d))

	for i, num := range d {
		a := (2 * c * num) / h
		q := math.Sqrt(float64(a))

		var b int = int(math.Round(q))

		s[i] = b
	}

	return s
}

