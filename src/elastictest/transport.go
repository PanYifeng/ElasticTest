package elastictest

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func EncodeRequest(c context.Context, r *http.Request, request interface{}) error {
	return encodeRequest(c, r, request)
}

func DecodeCreateIndexRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request createIndexRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeIndexByStructRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request indexByStructRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	return encodeResponse(c, w, response)
}

func DecodeCreateIndexResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response createIndexResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeIndexByStructResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response indexByStructResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeGetByIDResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response getByIDResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
