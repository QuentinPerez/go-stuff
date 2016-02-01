package main

import (
	"fmt"
	"reflect"

	"github.com/Sirupsen/logrus"
	"github.com/olivere/elastic"
)

type Data struct {
	StringValue string `json:"StringValue"`
	IntValue    int    `json:"IntValue"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.99.100:9200"), elastic.SetSniff(false))
	if err != nil {
		logrus.Fatal(err)
	}
	info, code, err := client.Ping().URL("http://192.168.99.100:9200").Do()
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("Elasticsearch returned with code %v and version %v\n", code, info.Version.Number)
	exist, err := client.IndexExists("try").Do()
	if err != nil {
		logrus.Fatal(err)
	}
	if !exist {
		_, err := client.CreateIndex("try").Do()
		if err != nil {
			logrus.Fatal(err)
		}
	}
	d := Data{"coucou", 1}
	put1, err := client.Index().Index("try").Type("Data").Id("1").BodyJson(d).Do()
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	d = Data{"foobar", 2}
	put2, err := client.Index().Index("try").Type("Data").Id("2").BodyJson(d).Do()
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)

	// Flush to make sure the documents got written.
	_, err = client.Flush().Index("twitter").Do()
	if err != nil {
		panic(err)
	}
	get1, err := client.Get().
		Index("try").
		Type("Data").
		Id("1").
		Do()
	if err != nil {
		logrus.Fatal(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}

	termQuery := elastic.NewTermQuery("StringValue", "coucou")
	searchResult, err := client.Search().
		Index("try").              // search in index "twitter"
		Query(&termQuery).         // specify the query
		Sort("StringValue", true). // sort by "user" field, ascending
		From(0).Size(10).          // take documents 0-9
		Pretty(true).              // pretty print request and response JSON
		Do()                       // execute
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	var ttyp Data
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(Data); ok {
			fmt.Printf("  Tweet by %s: %d\n", t.StringValue, t.IntValue)
		}
	}
	fmt.Printf("Found a total of %d data\n", searchResult.TotalHits())
	_, err = client.DeleteIndex("try").Do()
	if err != nil {
		logrus.Fatal(err)
	}
}
