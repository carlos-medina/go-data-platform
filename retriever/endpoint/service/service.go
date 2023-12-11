package service

import (
	"github.com/arquivei/foundationkit/errors"
)

type Response struct {
	UserID  int
	DataID  int
	Version int
	Content string
}

type Request struct {
	DataId int
}

type Service interface {
	Run(r Request) (Response, error)
}

type RetrieverService struct {
	MySQL MySQLGateway
}

func (s *RetrieverService) Run(r Request) (Response, error) {
	const op = errors.Op("service.RetrieverService.Run")

	response, err := s.MySQL.GetByDataId(r.DataId)

	if err != nil {
		return Response{}, errors.E(op, err)
	}

	return response, nil
}
