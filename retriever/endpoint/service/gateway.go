package service

type MySQLGateway interface {
	GetByDataId(dataId int) (Response, error)
}
