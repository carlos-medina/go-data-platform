package main

import (
	"fmt"

	"github.com/carlos-medina/go-data-platform/endpoint"
	"github.com/carlos-medina/go-data-platform/endpoint/gateway"

	"github.com/arquivei/foundationkit/errors"
)

// This Service was made just to test the new adapter
// The correct one will be much simpler, will have its own record struct
// and the service will be used by the endpoint
// Also, it will be an interface and will have a struct that implements
// its methods
func Service(record endpoint.Record, mySQL *gateway.MySQLAdapter) error {
	const op = errors.Op("main.service")

	// savedRecord, err := mySQL.Get(record.DataID)

	// if err != nil && err != sql.ErrNoRows {
	// 	errors.E(op, err)
	// }

	// // first time inserting
	// if err != nil && err == sql.ErrNoRows {
	// 	err = insertNewRow(record, mySQL)
	// 	if err != nil {
	// 		return errors.E(op, err)
	// 	}
	// } else {
	// 	// have to replace existing record
	// 	// Will be implemented with the update command

	// }

	err := insertNewRow(record, mySQL)

	if err != nil {
		return errors.E(op, err)
	}

	return nil
}

func insertNewRow(record endpoint.Record, mySQL *gateway.MySQLAdapter) error {
	const op = errors.Op("main.insertNewRow")

	fmt.Println("kkkkkk1")

	err := mySQL.Insert(record)

	if err != nil {
		return errors.E(op, err)
	}

	return nil
}

// func insertExistingRow()
