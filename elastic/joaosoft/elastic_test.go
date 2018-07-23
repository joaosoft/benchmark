package elastic_joaosoft

import (
	"testing"

	"fmt"
	"strconv"
	"time"

	"net/http"
	"os"

	structs "benchmark/elastic"

	"sync"

	"github.com/joaosoft/elastic"
	log "github.com/joaosoft/logger"
)

var client = elastic.NewElastic()

func BenchmarkJoaosoftElastic(b *testing.B) {
	// index create with mapping
	createIndexWithMapping()

	// document create
	createDocumentWithId("1")
	createDocumentWithId("2")
	generatedId := createDocumentWithoutId()

	// document update
	updateDocumentWithId("1")
	updateDocumentWithId("2")

	// document search
	// wait elastic to index the last update...
	<-time.After(time.Second * 2)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			searchDocument("luis")
			wg.Done()
		}()

	}
	wg.Wait()

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
	id, err := client.Create().Index("persons").Type("person").Id(id).Body(structs.Person{
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
	id, err := client.Create().Index("persons").Type("person").Body(structs.Person{
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
	id, err := client.Create().Index("persons").Type("person").Id(id).Body(structs.Person{
		Name: "luis",
		Age:  age + 20,
	}).Execute()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\nupdated person with id %s\n", id)
	}
}

func searchDocument(name string) {
	var data []structs.Person

	d1 := elastic.TemplateData{Data: map[string]interface{}{"name": name}}

	// document search
	dir, _ := os.Getwd()
	err := client.Search().
		Index("persons").
		Type("person").
		Object(&data).
		Template(dir+"/examples/templates", "get.example.1.template", &d1, false).
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
