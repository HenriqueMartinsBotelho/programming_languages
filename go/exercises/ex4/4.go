package ex4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func ConvertToSlice(input string) ([]int, error) {
    input = strings.TrimSpace(input) 
    if input == "" {
        return []int{}, nil
    }

    numberStrings := strings.Split(input, ",")
    var numbers []int
    for _, numberString := range numberStrings {
        number, err := strconv.Atoi(strings.TrimSpace(numberString))
        if err != nil {
            return nil, err
        }
        numbers = append(numbers, number)
    }
    return numbers, nil
}


func Ex4() {
	fmt.Println("Enter a sequence of comma-separated numbers:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) 

	numbers, err := ConvertToSlice(input)
	if err != nil {
		fmt.Println("Error converting string to number:", err)
		return
	}

	fmt.Println("Generated slice:", numbers)
}

// Write a program which accepts a sequence of comma-separated numbers from console and generate an slice out of them. Return the slice.

// Suppose the following input is supplied to the program: 34, 67, 55, 33, 12, 98.

// Then, the output should be: [34 67 55 33 12 98]
