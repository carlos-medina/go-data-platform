package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/carlos-medina/go-data-platform/retriever/endpoint/gateway"

	"github.com/arquivei/foundationkit/errors"
	"github.com/gorilla/mux"
)

func main() {
	mySQLAdapter := MustNewMySQLAdapter()

	r := mux.NewRouter()
	r.HandleFunc("/records", getHandleRecords(mySQLAdapter)).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func getHandleRecords(mySQLAdapter *gateway.MySQLAdapter) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		const op = errors.Op("main.getHandleRecords")
		dataIdStr := r.URL.Query().Get("data_id")

		// TODO: prints on screen that data filter must be provided
		if dataIdStr == "" {
			panic(errors.E(op, "filter data_id not provided"))
		}

		dataId, err := strconv.Atoi(dataIdStr)

		// TODO: prints on screen that data must be a valid integer
		if err != nil {
			panic(errors.E(op, err))
		}

		record, err := mySQLAdapter.GetByDataId(dataId)

		// TODO: prints on screen that record was not found
		if err != nil {
			panic(errors.E(op, err))
		}

		response, err := json.Marshal(record)

		// TODO: prints on screen that response could not be encoded
		if err != nil {
			panic(errors.E(op, err))
		}

		_, err = w.Write(response)

		if err != nil {
			panic(errors.E(op, err))
		}
	}

	return fn
}
