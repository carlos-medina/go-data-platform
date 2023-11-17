package gateway

import (
	"github.com/carlos-medina/go-data-platform/retriever/endpoint"
)

type MySQLGateway interface {
	GetByDataId(dataIn int) endpoint.Record
}
