package elastic_joaosoft

import (
	"github.com/joaosoft/elastic"
	"testing"

	"fmt"
	"strconv"
	"time"

	"os"

	"sync"

	"github.com/joaosoft/elasticsearch"
	log "github.com/joaosoft/logger"
)

var client, _ = elasticsearch.NewElastic()

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
	_, err := client.Index().Index("persons").Body([]byte(`
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
`)).Create()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ncreated mapping for persons index\n")
	}
}

func createDocumentWithId(id string) {
	// document create with id
	age, _ := strconv.Atoi(id)
	_, err := client.Document().Index("persons").Type("person").Id(id).Body(Person{
		Name: "joao",
		Age:  age + 20,
	}).Create()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ncreated a new person with id %s\n", id)
	}
}

func createDocumentWithoutId() string {
	// document create without id
	response, err := client.Document().Index("persons").Type("person").Body(Person{
		Name: "joao",
		Age:  30,
	}).Create()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ncreated a new person with id %s\n", response.ID)
	}

	return response.ID
}

func updateDocumentWithId(id string) {
	// document update with id
	age, _ := strconv.Atoi(id)
	response, err := client.Document().Index("persons").Type("person").Id(id).Body(Person{
		Name: "luis",
		Age:  age,
	}).Update()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\nupdated person with id %s\n", response.ID)
	}
}

func searchDocument(name string) {
	var data []Person

	d1 := elastic.SearchTemplate{Data: map[string]interface{}{"name": name, "size": "10000"}}

	// document search
	dir, _ := os.Getwd()
	response, err := client.Search().
		Index("persons").
		Type("person").
		Object(&data).
		Template(dir+"/examples/templates", "get.example.search.template", &d1, false).
		Search()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\nsearch person by name:%s %+v\n", name, response.Hits.Hits)
	}
}

func deleteDocumentWithId(id string) {
	response, err := client.Document().Index("persons").Type("person").Id(id).Delete()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ndeleted person with id %s\n", response.ID)
	}
}

func existsIndex(index string) {
	status, err := client.Index().Index(index).Exists()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\nexists index? %t\n", status)
	}
}

func deleteIndex() {
	_, err := client.Index().Index("persons").Delete()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ndeleted persons index\n")
	}
}
