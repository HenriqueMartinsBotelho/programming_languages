// Write a program which takes 2 digits, X,Y as input and generates a 2-dimensional array. The element value in the i-th row and j-th column of the array should be i*j. Note: i=0,1.., X-1; j=0,1,¡­Y-1.

// Example

// Suppose the following inputs are given to the program: 3,5

// Then, the output of the program should be: [[0, 0, 0, 0, 0], [0, 1, 2, 3, 4], [0, 2, 4, 6, 8]]

package ex7

func Ex7(X, Y int) [][]int {
    m := make([][]int, X)
    for i := 0; i < X; i++ {
        m[i] = make([]int, Y)
        for j := 0; j < Y; j++ {
            m[i][j] = i * j
        }
    }

    return m
}

// solução alternativa
func Ex7_2(x, y int) [][]int {
    slice := [][]int{}

    for i := 0; i < x; i++ {
        row := []int{}
        for j := 0; j < y; j++ {
            p := i * j
            row = append(row, p)
        }
        slice = append(slice, row)
    }
    return slice
}
