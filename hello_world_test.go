package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}

func BenchmarkTable(b *testing.B) {
	bench := []struct {
		name    string
		request string
	}{
		{
			name:    "Eko",
			request: "Eko",
		},
		{
			name:    "Dono",
			request: "Dono",
		},
		{
			name:    "DonoWardhana",
			request: "DonoWardhana",
		},
	}

	for _, item := range bench {
		b.Run(item.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(item.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {

	b.Run("Eko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Eko")
		}
	})

	b.Run("Dono", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Dono")
		}
	})
}

func BenchmarkHelloWorld2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kurniawan")
	}
}

func TestHelloWorld1(t *testing.T) {
	res := HelloWorld("Eko")

	assert.Equal(t, "Hello Eko", res)
	fmt.Print("sudah di eksekusi")
}
