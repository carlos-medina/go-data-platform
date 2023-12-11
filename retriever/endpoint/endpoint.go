package endpoint

import (
	"github.com/carlos-medina/go-data-platform/retriever/endpoint/service"

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

type Endpoint interface {
	Run(r Request) (Response, error)
}

type RetrieverEndpoint struct {
	Service service.Service
}

func (re *RetrieverEndpoint) Run(r Request) (Response, error) {
	const op = errors.Op("endpoint.RetrieverEndpoint.Run")

	serviceRequest := service.Request{
		DataId: r.DataId,
	}

	serviceResponse, err := re.Service.Run(serviceRequest)

	if err != nil {
		return Response{}, errors.E(op, err)
	}

	return Response{
		UserID:  serviceResponse.UserID,
		DataID:  serviceResponse.DataID,
		Version: serviceResponse.Version,
		Content: serviceResponse.Content,
	}, nil
}
