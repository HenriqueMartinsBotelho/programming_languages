package ex2

func Factorial(number int) int {
	if number == 1{
		return 1
	}
	
	return number * Factorial(number - 1)
}

// Exercise Description

// Write a program which can compute the factorial of a given numbers. The results should be printed in a comma-separated sequence on a single line.

// Suppose the following input is supplied to the program: 8

// Then, the output should be: 40320