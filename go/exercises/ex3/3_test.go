package ex3

import "testing"


func mapsAreEqual(a, b map[int]int) bool {
    if len(a) != len(b) {
        return false
    }

    for k, v := range a {
        if w, ok := b[k]; !ok || v != w {
            return false
        }
    }

    return true
}


func TestEx3(t *testing.T){
	m := map[int]int{
		1: 1,
		2: 4,
		3: 9,
		4: 16,
		5: 25,
	}

	want := m
	got := Ex3(5)

	if !mapsAreEqual(got, want) {
        t.Errorf("Ex3 = got %v, want %v", got, want)
    }

}



// With a given integral number n, write a program to generate a map that contains (i, i*i) such that is an integral number between 1 and n (both included), and then the program should print the map with representation of the value

// Suppose the following input is supplied to the program: 8

// Then, the output should be: map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]
