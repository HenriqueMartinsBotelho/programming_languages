package ex2

import "testing"


func TestEx2(t *testing.T){
	want := 120
	got := Factorial(5)

	if got != want {
		t.Errorf("Ex2 = %v, want %v", got, want)
	}
}
