package service

import (
	"fmt"
	"strings"

	"github.com/carlos-medina/go-data-platform/ingestor/endpoint"
	"github.com/carlos-medina/go-data-platform/ingestor/endpoint/gateway"

	"github.com/arquivei/foundationkit/errors"
)

type Service interface {
	Run(record endpoint.Record) error
}

type IService struct {
	MySQL gateway.MySQLGateway
}

func (s *IService) Run(record endpoint.Record) error {
	const op = errors.Op("service.IService.Run")

	currentRecord, err := s.MySQL.Get(record.DataID)

	if err != nil && !isNoRowsError(err) {
		return errors.E(op, err)
	}

	if err != nil {
		err = s.insert(record)
	} else {
		if currentRecord.Version >= record.Version {
			return errors.E(op, fmt.Sprintf("Can't update current version because new one is equal or lower - current version: %v; new version: %v", currentRecord.Version, record.Version))
		}

		err = s.update(record)
	}

	if err != nil {
		return errors.E(op, err)
	}

	return nil
}

func (s *IService) insert(record endpoint.Record) error {
	const op = errors.Op("service.IService.insert")

	err := s.MySQL.Insert(record)

	if err != nil {
		return errors.E(op, err)
	}

	return nil
}

func (s *IService) update(record endpoint.Record) error {
	const op = errors.Op("service.IService.update")

	err := s.MySQL.Update(record)

	if err != nil {
		return errors.E(op, err)
	}

	return nil
}

func isNoRowsError(err error) bool {
	sError := err.Error()

	return strings.Contains(sError, "sql: no rows in result set")
}
