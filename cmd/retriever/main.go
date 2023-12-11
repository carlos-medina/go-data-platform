package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/carlos-medina/go-data-platform/retriever/endpoint"

	"github.com/arquivei/foundationkit/errors"
	"github.com/gorilla/mux"
)

type Response struct {
	UserID  int    `json:"user_id"`
	DataID  int    `json:"data_id"`
	Version int    `json:"version"`
	Content string `json:"content"`
}

func main() {
	endpoint := MustNewEndpoint()

	r := mux.NewRouter()
	r.HandleFunc("/records", getHandleRecords(endpoint)).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func getHandleRecords(endpoint endpoint.Endpoint) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		const op = errors.Op("main.getHandleRecords")

		endpointRequest, err := translateRequest(r)

		// TODO: print on screen a valid error for the user
		if err != nil {
			panic(errors.E(op, err))
		}

		endpointResponse, err := endpoint.Run(endpointRequest)

		// TODO: print on screen a valid error for the user
		if err != nil {
			panic(errors.E(op, err))
		}

		response := translateResponse(endpointResponse)

		jsonResponse, err := json.Marshal(response)

		// TODO: print on screen a valid error for the user
		if err != nil {
			panic(errors.E(op, err))
		}

		_, err = w.Write(jsonResponse)

		// TODO: print on screen a valid error for the userS
		if err != nil {
			panic(errors.E(op, err))
		}
	}

	return fn
}

func translateRequest(r *http.Request) (endpoint.Request, error) {
	const op = errors.Op("main.translateRequest")
	dataIdStr := r.URL.Query().Get("data_id")

	if dataIdStr == "" {
		return endpoint.Request{}, errors.E(op, "Filter data_id must be provided")
	}

	dataId, err := strconv.Atoi(dataIdStr)

	if err != nil {
		return endpoint.Request{}, errors.E(op, err)
	}

	return endpoint.Request{
		DataId: dataId,
	}, nil
}

func translateResponse(endpointResponse endpoint.Response) Response {
	return Response{
		UserID:  endpointResponse.UserID,
		DataID:  endpointResponse.DataID,
		Version: endpointResponse.Version,
		Content: endpointResponse.Content,
	}
}
