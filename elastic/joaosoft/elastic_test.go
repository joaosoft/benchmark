package elastic_joaosoft

import (
	"testing"

	"fmt"
	"strconv"
	"time"

	"net/http"
	"os"

	"sync"

	"github.com/joaosoft/elastic"
	log "github.com/joaosoft/logger"
)

var client = elastic.NewElastic()

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func BenchmarkJoaosoftElastic(b *testing.B) {
	// index create with mapping
	createIndexWithMapping()

	// document create
	for i := 1; i <= 10000; i++ {
		createDocumentWithId(strconv.Itoa(i))
	}
	generatedId := createDocumentWithoutId()

	// document update
	updateDocumentWithId("90009")
	updateDocumentWithId("90008")

	// document search
	// wait elastic to index the last update...
	<-time.After(time.Second * 2)

	var wg1 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg1.Add(1)
		go func() {
			searchDocument("luis")
			wg1.Done()
		}()
	}
	wg1.Wait()

	var wg2 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			searchDocument("joao")
			wg2.Done()
		}()
	}
	wg2.Wait()

	return

	// document delete
	deleteDocumentWithId(generatedId)

	// index exists
	existsIndex("persons")
	existsIndex("bananas")

	// index delete
	deleteIndex()
}

func createIndexWithMapping() {
	err := client.CreateIndex().Index("persons").Body([]byte(`
{
  "mappings": {
    "person": {
      "properties": {
        "age": {
          "type": "long"
        },
        "name": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        }
      }
    }
  }
}
`)).Execute()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ncreated mapping for persons index\n")
	}
}

func createDocumentWithId(id string) {
	// document create with id
	age, _ := strconv.Atoi(id)
	id, err := client.Create().Index("persons").Type("person").Id(id).Body(Person{
		Name: "joao",
		Age:  age + 20,
	}).Execute()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ncreated a new person with id %s\n", id)
	}
}

func createDocumentWithoutId() string {
	// document create without id
	id, err := client.Create().Index("persons").Type("person").Body(Person{
		Name: "joao",
		Age:  30,
	}).Execute()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ncreated a new person with id %s\n", id)
	}

	return id
}

func updateDocumentWithId(id string) {
	// document update with id
	age, _ := strconv.Atoi(id)
	id, err := client.Create().Index("persons").Type("person").Id(id).Body(Person{
		Name: "luis",
		Age:  age,
	}).Execute()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\nupdated person with id %s\n", id)
	}
}

func searchDocument(name string) {
	var data []Person

	d1 := elastic.SearchTemplate{Data: map[string]interface{}{"name": name, "size": "10000"}}

	// document search
	dir, _ := os.Getwd()
	err := client.Search().
		Index("persons").
		Type("person").
		Object(&data).
		Template(dir+"/examples/templates", "get.example.search.template", &d1, false).
		Execute()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\nsearch person by name:%s %+v\n", name, data)
	}
}

func deleteDocumentWithId(id string) {
	err := client.Delete().Index("persons").Type("person").Id(id).Execute()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ndeleted person with id %s\n", "1")
	}
}

func existsIndex(index string) {
	status, err := client.ExistsIndex().Index(index).Execute()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\nexists index? %t\n", status == http.StatusOK)
	}
}

func deleteIndex() {
	err := client.DeleteIndex().Index("persons").Execute()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ndeleted persons index\n")
	}
}
