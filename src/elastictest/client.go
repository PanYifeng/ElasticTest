// client for the demo service.
package elastictest

import (
	"net/url"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewCreateIndexCli(host string) Service {
	url, _ := url.Parse("http://" + host + "/server/createIndex")
	createIndexServer := httptransport.NewClient(
		"POST",
		url,
		EncodeRequest,
		DecodeCreateIndexResponse,
	).Endpoint()

	return ElasticEndpoints{
		CreateIndexEndpoint: createIndexServer,
	}
}

func NewIndexByStructCli(host string) Service {
	url, _ := url.Parse("http://" + host + "/server/indexByStruct")
	indexByStructServer := httptransport.NewClient(
		"POST",
		url,
		EncodeRequest,
		DecodeIndexByStructResponse,
	).Endpoint()

	return ElasticEndpoints{
		IndexByStructEndpoint: indexByStructServer,
	}
}

func NewGetByIDCli(host string) Service {
	url, _ := url.Parse("http://" + host + "/server/getByID")
	getByIDServer := httptransport.NewClient(
		"POST",
		url,
		EncodeRequest,
		DecodeGetByIDResponse,
	).Endpoint()

	return ElasticEndpoints{
		GetByIDEndpoint: getByIDServer,
	}
}
