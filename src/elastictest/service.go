package elastictest

import (
	"context"
	"errors"
	"log"
	"os"

	elastic "gopkg.in/olivere/elastic.v5"
)

type Service interface {
	CreateIndex(ctx context.Context, indexName string) error
	IndexByStruct(ctx context.Context, indexName string, typeName string, id string, data interface{}) (*elastic.IndexResponse, error)
	GetByID(ctx context.Context, indexName string, typeName string, id string) (*elastic.GetResult, error)
}

type ElasticSvc struct{}

var (
	client, _          = elastic.NewClient(elastic.SetTraceLog(log.New(os.Stderr, "[[ELASTIC]]", 0)))
	ErrNotAcknowledged = errors.New("Not acknowledged")
)

func (ElasticSvc) CreateIndex(ctx context.Context, indexName string) error {
	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		// Handle error
		// panic(err)
		return err
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex(indexName).Do(ctx)
		if err != nil {
			// Handle error
			//panic(err)
			return err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			return ErrNotAcknowledged
		}
	}
	return nil
}

func (ElasticSvc) IndexByStruct(ctx context.Context, indexName string, typeName string, id string, data interface{}) (*elastic.IndexResponse, error) {
	put, err := client.Index().
		Index(indexName).
		Type(typeName).
		Id(id).
		BodyJson(data).
		Do(ctx)
	return put, err
}

func (ElasticSvc) GetByID(ctx context.Context, indexName string, typeName string, id string) (*elastic.GetResult, error) {
	get, err := client.Get().
		Index(indexName).
		Type(typeName).
		Id(id).
		Do(ctx)
	return get, err
}
