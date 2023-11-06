package gateway

import (
	"github.com/carlos-medina/go-data-platform/endpoint"
)

type MySQLGateway interface {
	Get(dataId int) (endpoint.Record, error)
	Insert(endpoint.Record) error
}
