package iteracao

import "testing"

func BenchmarkRepeticao(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repetir("")
	}
}