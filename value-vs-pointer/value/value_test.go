package value

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

func BenchmarkValue(b *testing.B) {
	b.ReportAllocs()
	data := Data{
		A: "AAAAAAAAAAAAAAAAAAA",
		B: "BBBBBBBBBBBBBBBBBBB",
		C: "CCCCCCCCCCCCCCCCCCC",
		D: "DDDDDDDDDDDDDDDDDDD",
		E: "EEEEEEEEEEEEEEEEEEE",
	}
	data.F = data

	for i := 0; i < b.N; i++ {
		dummyFunc(data)
	}
}

func dummyFunc(data Data) error {
	if data.A == data.B {
		panic("fail")
	}
	return nil
}
