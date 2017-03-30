package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"elastictest"
	elastic "gopkg.in/olivere/elastic.v5"
)

// Tweet is a structure used for serializing/deserializing data in Elasticsearch.
type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

func main() {
	log.Println("start to create index 'test'")
	createIndexCli := elastictest.NewCreateIndexCli("localhost:9000")
	err := createIndexCli.CreateIndex(context.Background(), "test")
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(s)

	log.Println("start to index tweet1")
	indexByStructCli := elastictest.NewIndexByStructCli("localhost:9000")
	tweet1 := Tweet{User: "olivere", Message: "Take Five", Retweets: 0}
	put, err := indexByStructCli.IndexByStruct(context.Background(), "test", "tweet", "1", tweet1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Indexed %+v %s to index %s, type %s\n", tweet1, put.Id, put.Index, put.Type)

	log.Println("start to get ID1")
	getByIDCli := elastictest.NewGetByIDCli("localhost:9000")
	get, err := getByIDCli.GetByID(context.Background(), "test", "tweet", "1")
	if err != nil {
		log.Fatal(err)
	}
	var tweetGot Tweet
	err = json.Unmarshal(*get.Source, &tweetGot)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Got document from index 'test', type 'tweet': %v -- User:%v -- Message:%v\n", tweetGot, tweetGot.User, tweetGot.Message)
}
