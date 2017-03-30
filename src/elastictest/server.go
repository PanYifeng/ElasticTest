package elastictest

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func Run() {
	svc := ElasticSvc{}

	createIndexHandler := httptransport.NewServer(
		MakeCreateIndexEndpoint(svc),
		DecodeCreateIndexRequest,
		EncodeResponse,
	)

	indexByStructHandler := httptransport.NewServer(
		MakeIndexByStructEndpoint(svc),
		DecodeIndexByStructRequest,
		EncodeResponse,
	)

	getByIDHandler := httptransport.NewServer(
		MakeGetByIDEndpoint(svc),
		DecodeGetByIDRequest,
		EncodeResponse,
	)

	http.Handle("/server/createIndex", createIndexHandler)
	http.Handle("/server/indexByStruct", indexByStructHandler)
	http.Handle("/server/getByID", getByIDHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
