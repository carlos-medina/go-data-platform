package service

import (
	"github.com/carlos-medina/go-data-platform/endpoint"
	"github.com/carlos-medina/go-data-platform/endpoint/gateway"

	"github.com/arquivei/foundationkit/errors"
)

type Service interface {
	Run(record endpoint.Record) error
}

type IService struct {
	MySQL *gateway.MySQLAdapter
}

func (s *IService) Run(record endpoint.Record) error {
	const op = errors.Op("service.IService.Run")

	err := s.MySQL.Insert(record)

	if err != nil {
		return errors.E(op, err)
	}

	return nil
}
