// Write a program which accepts a sequence of comma separated 4 digit binary numbers as its input and then check whether they are divisible by 5 or not. The numbers that are divisible by 5 are to be printed in a comma separated sequence.

// Example: 0100,0011,1010,1001

// Then the output should be: 1010

package ex8

import "strconv"

func Ex8(sequence []string) []string {

	res := []string{}

	for _, s := range sequence {
		converted, err := strconv.ParseInt(s,  2,  64);
		if err != nil {
			println("Erro:", err)
		}
		
		if converted % 5 == 0 {
			res = append(res, s)
		}

	}

	return res

}
