package ex3


func Ex3(number int) map[int]int{
	
	m := make(map[int]int)

	for i := 1; i <= number; i++ {
		m[i] = i*i
	}

	return m
}