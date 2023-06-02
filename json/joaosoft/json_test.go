package elastic_joaosoft

import (
	"fmt"
	"github.com/joaosoft/json"
	"testing"
)

type address struct {
	Ports  []int             `db.read:"ports" json:"ports"`
	Street string            `db.read:"street" json:"street"`
	Number float64           `db:"number" db.write:"number" json:"number"`
	Map    map[string]string `db:"map" json:"map"`
}

type person struct {
	Name      string              `db:"name" json:"name"`
	Age       int                 `db:"age" json:"age"`
	Address   *address            `db:"address" json:"address"`
	Numbers   []int               `db:"numbers" json:"numbers"`
	Others    map[string]*address `db:"others" json:"others"`
	Addresses []*address          `db:"addresses" json:"addresses"`
}

func BenchmarkJsonMarshal(b *testing.B) {
	addr := &address{
		Street: "street one",
		Number: 1.2,
		Map:    map[string]string{`"ola" "joao"`: `"adeus" "joao"`, "c": "d"},
	}

	example := person{
		Name:      "joao",
		Age:       30,
		Address:   addr,
		Numbers:   []int{1, 2, 3},
		Others:    map[string]*address{`"ola" "joao"`: addr, "c": addr},
		Addresses: []*address{addr, addr},
	}

	// with tags "db" and "db.read"
	// marshal
	bytes, err := json.Marshal(example, "db", "db.read")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	addr := &address{
		Street: "street one",
		Number: 1.2,
		Map:    map[string]string{`"ola" "joao"`: `"adeus" "joao"`, "c": "d"},
	}

	example := person{
		Name:      "joao",
		Age:       30,
		Address:   addr,
		Numbers:   []int{1, 2, 3},
		Others:    map[string]*address{`"ola" "joao"`: addr, "c": addr},
		Addresses: []*address{addr, addr},
	}

	// with tags "db" and "db.read"
	// marshal
	bytes, err := json.Marshal(example, "db", "db.read")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

	// unmarshal
	var newExample person
	err = json.Unmarshal(bytes, &newExample, "db", "db.read")
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n:: Example: %+v", newExample)
	fmt.Printf("\n:: Address: %+v\n\n\n", newExample.Address)
}
