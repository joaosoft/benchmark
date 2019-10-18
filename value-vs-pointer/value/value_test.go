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
	G interface{}
	H interface{}
	I interface{}
	J interface{}
	K interface{}
	L interface{}
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
	data.G = data
	data.H = data
	data.I = data
	data.J = data
	data.K = data
	data.L = data

	for i := 0; i < 5; i++ {
		dummyFunc(data, 10)
	}
}

func dummyFunc(data Data, step int) error {
	if data.A == data.B {
		panic("fail")
	}

	if step > 0 {
		return dummyFunc(data, step-1)
	} else {
		return nil
	}
}
