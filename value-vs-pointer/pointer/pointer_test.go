package pointer

import (
	"testing"
)

type Data struct {
	A string
	B string
	C string
	D string
	E string
	F interface{}
}

func BenchmarkPointer(b *testing.B) {
	data := Data{
		A: "AAAAAAAAAAAAAAAAAAA",
		B: "BBBBBBBBBBBBBBBBBBB",
		C: "CCCCCCCCCCCCCCCCCCC",
		D: "DDDDDDDDDDDDDDDDDDD",
		E: "EEEEEEEEEEEEEEEEEEE",
	}
	data.F = &data

	for i := 0; i < b.N; i++ {
		dummyFunc(&data)
	}
}

func dummyFunc(data *Data) error {
	return nil
}
