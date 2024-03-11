// Write a program, which will find all such numbers between 100 and 300 (both included) such that each digit of the number is an even number.

package ex9

func Ex9() []int {

	res := []int{}

	for i := 100; i <= 300 ; i++ {
		
		firstDigit := i %  10
        secondDigit := (i /  10) %  10
        thirdDigit := i /  100

		if firstDigit % 2 ==  0 && secondDigit % 2 ==  0 && thirdDigit % 2 ==  0 {
            res = append(res, i)
        }
	}


	return res
}


