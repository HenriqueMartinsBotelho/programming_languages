package ex7

import (
	"reflect"
	"testing"
)

func TestEx7(t *testing.T) {
    tests := []struct {
        X        int
        Y        int
        expected [][]int
    }{
        {
            X: 2,
            Y: 2,
            expected: [][]int{
                {0, 0},
                {0, 1},
            },
        },
        {
            X: 3,
            Y: 4,
            expected: [][]int{
                {0, 0, 0, 0},
                {0, 1, 2, 3},
                {0, 2, 4, 6},
            },
        },
    }

    for _, tt := range tests {
        result := Ex7(tt.X, tt.Y)
        if !reflect.DeepEqual(result, tt.expected) {
            t.Errorf("Ex7(%d, %d) got %v, want %v", tt.X, tt.Y, result, tt.expected)
        }
    }
}
