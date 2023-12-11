package main

import (
	"database/sql"

	"github.com/carlos-medina/go-data-platform/retriever/endpoint"
	"github.com/carlos-medina/go-data-platform/retriever/endpoint/service"

	"github.com/arquivei/foundationkit/errors"
	"github.com/go-sql-driver/mysql"
)

func MustNewMySQLAdapter() *service.MySQLAdapter {
	const op = errors.Op("main.MustNewMySQLAdapter")

	cfg := mysql.Config{
		User:   "root",
		Passwd: "admin",
		Net:    "tcp",
		Addr:   "172.18.0.2:3306",
		DBName: "go_data_platform",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(errors.E(op, err))
	}

	err = db.Ping()

	if err != nil {
		panic(errors.E(op, err))
	}

	return &service.MySQLAdapter{
		DB:    db,
		Table: "records",
	}
}

func MustNewService() *service.RetrieverService {
	mySQLAdapter := MustNewMySQLAdapter()

	return &service.RetrieverService{
		MySQL: mySQLAdapter,
	}
}

func MustNewEndpoint() *endpoint.RetrieverEndpoint {
	service := MustNewService()

	return &endpoint.RetrieverEndpoint{
		Service: service,
	}
}
