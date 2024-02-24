// Write a program that calculates and prints the value according to the given formula:

// Q = Square root of [(2 * C * D)/H]

// Following are the fixed values of C and H:

//     C is 50. H is 30.

//     D is the variable whose values should be input to your program in a comma-separated sequence.

package ex6

import (
	"reflect"
	"testing"
)

func TestEx006(t *testing.T) {

	input := []int{100, 150, 180}
	want := []int{18, 22, 24}
	got := Ex006(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ex006() = %v, want %v", got, want)
	}
}