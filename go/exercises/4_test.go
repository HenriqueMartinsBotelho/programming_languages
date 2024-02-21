package exercises

import (
	"reflect"
	"testing"
)

func TestConvertToSlice(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{"1,2,3", []int{1, 2, 3}},
		{"4, 5, 6", []int{4, 5, 6}},
		{"", []int{}},
	}

	for _, tt := range tests {
		got, err := ConvertToSlice(tt.input)
		if err != nil {
			t.Errorf("ConvertToSlice(%q) returned an error: %v", tt.input, err)
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ConvertToSlice(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}