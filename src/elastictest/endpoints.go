package elastictest

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	elastic "gopkg.in/olivere/elastic.v5"
)

type ElasticEndpoints struct {
	CreateIndexEndpoint   endpoint.Endpoint
	IndexByStructEndpoint endpoint.Endpoint
	GetByIDEndpoint       endpoint.Endpoint
}

// for clients
func (e ElasticEndpoints) CreateIndex(ctx context.Context, indexName string) error {
	request := createIndexRequest{IndexName: indexName}
	response, err := e.CreateIndexEndpoint(ctx, request)
	if err != nil {
		return err
	}
	return response.(createIndexResponse).Err
}

func (e ElasticEndpoints) IndexByStruct(ctx context.Context, indexName string, typeName string, id string, data interface{}) (*elastic.IndexResponse, error) {
	request := indexByStructRequest{IndexName: indexName, TypeName: typeName, Id: id, Data: data}
	response, err := e.IndexByStructEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(indexByStructResponse).Put, response.(indexByStructResponse).Err
}

func (e ElasticEndpoints) GetByID(ctx context.Context, indexName string, typeName string, id string) (*elastic.GetResult, error) {
	request := getByIDRequest{IndexName: indexName, TypeName: typeName, Id: id}
	response, err := e.GetByIDEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(getByIDResponse).Get, response.(getByIDResponse).Err
}

// for server
func MakeCreateIndexEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		createIndexReq := request.(createIndexRequest)
		err := s.CreateIndex(ctx, createIndexReq.IndexName)
		return createIndexResponse{
			Err: err,
		}, nil
	}
}

func MakeIndexByStructEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		indexByStructReq := request.(indexByStructRequest)
		put, err := s.IndexByStruct(ctx, indexByStructReq.IndexName, indexByStructReq.TypeName, indexByStructReq.Id, indexByStructReq.Data)
		return indexByStructResponse{
			Put: put,
			Err: err,
		}, nil
	}
}

func MakeGetByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		getByIDReq := request.(getByIDRequest)
		get, err := s.GetByID(ctx, getByIDReq.IndexName, getByIDReq.TypeName, getByIDReq.Id)
		return getByIDResponse{
			Get: get,
			Err: err,
		}, nil
	}
}

type createIndexRequest struct{ IndexName string }

type createIndexResponse struct{ Err error }

type indexByStructRequest struct {
	IndexName string
	TypeName  string
	Id        string
	Data      interface{}
}

type indexByStructResponse struct {
	Put *elastic.IndexResponse
	Err error
}

type getByIDRequest struct {
	IndexName string
	TypeName  string
	Id        string
}

type getByIDResponse struct {
	Get *elastic.GetResult
	Err error
}
