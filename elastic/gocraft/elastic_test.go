package elastic_gocraft

import (
	"context"
	"testing"

	"fmt"
	"strconv"
	"time"

	structs "benchmark/elastic"

	"encoding/json"

	"sync"

	log "github.com/joaosoft/logger"
	elastic "github.com/olivere/elastic"
)

var client, _ = elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))

func BenchmarkGocraftElastic(b *testing.B) {
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
	_, err := client.CreateIndex("persons").Body(string([]byte(`
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
`))).Do(context.Background())

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ncreated mapping for persons index\n")
	}
}

func createDocumentWithId(id string) {
	// document create with id
	age, _ := strconv.Atoi(id)
	_, err := client.Index().Index("persons").Type("person").Id(id).BodyJson(structs.Person{
		Name: "joao",
		Age:  age + 20,
	}).Do(context.Background())

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ncreated a new person with id %s\n", id)
	}
}

func createDocumentWithoutId() string {
	// document create without id
	response, err := client.Index().Index("persons").Type("person").BodyJson(structs.Person{
		Name: "joao",
		Age:  30,
	}).Do(context.Background())

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ncreated a new person with id %s\n", response.Id)
	}

	return response.Id
}

func updateDocumentWithId(id string) {
	// document update with id
	age, _ := strconv.Atoi(id)
	_, err := client.Update().Index("persons").Type("person").Id(id).Doc(structs.Person{
		Name: "luis",
		Age:  age + 20,
	}).Do(context.Background())

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\nupdated person with id %s\n", id)
	}
}

func searchDocument(name string) {
	var data []structs.Person

	// document search
	result, err := client.Search().
		Index("persons").
		Type("person").
		Query(elastic.NewBoolQuery().Must(elastic.NewTermQuery("name", name))).
		Do(context.Background())

	if err == nil {
		if len(result.Hits.Hits) > 0 {
			p := structs.Person{}
			err = json.Unmarshal(*result.Hits.Hits[0].Source, &p)
			if err != nil {
				log.Error(err)
			}
			data = append(data, p)
		} else {
			log.Error(err)
		}
	}

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\nsearch person by name:%s %+v\n", name, data)
	}
}

func deleteDocumentWithId(id string) {
	_, err := client.Delete().Index("persons").Type("person").Id(id).Do(context.Background())

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ndeleted person with id %s\n", "1")
	}
}

func existsIndex(index string) {
	status, err := client.IndexExists(index).Do(context.Background())

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\nexists index? %t\n", status)
	}
}

func deleteIndex() {
	_, err := client.DeleteIndex("persons").Do(context.Background())

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ndeleted persons index\n")
	}
}
