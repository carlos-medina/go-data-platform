package gateway

import (
	"database/sql"
	"fmt"

	"github.com/carlos-medina/go-data-platform/endpoint"

	"github.com/arquivei/foundationkit/errors"
)

// mySQLDB interface is satisfied by sql.DB
type mySQLDB interface {
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
}

// MySQLAdapter implements MySQLGateway
// It uses mySQLDB interface that is satisfied by sql.DB
// and receives which table is being read or modified
type MySQLAdapter struct {
	DB    mySQLDB
	Table string
}

func (mySQL MySQLAdapter) Get(dataId int) (endpoint.Record, error) {
	const op = errors.Op("gateway.MySQLAdapter.Get")
	var record endpoint.Record

	row := mySQL.DB.QueryRow("SELECT user_id, data_id, version, content FROM ? WHERE data_id = ?", mySQL.Table, dataId)
	if err := row.Scan(&record.UserID, &record.DataID, &record.Version, &record.Content); err != nil {
		if err == sql.ErrNoRows {
			return record, errors.E(op, err)
		}
		return record, errors.E(op, err)
	}

	return record, nil
}

func (mySQL MySQLAdapter) Insert(r endpoint.Record) error {
	const op = errors.Op("gateway.MySQLAdapter.Insert")
	query := fmt.Sprintf("INSERT INTO %v (user_id, data_id, version, content) VALUES (?, ?, ?, ?)", mySQL.Table)

	fmt.Printf("endpoint.Record: %+v \n", r)
	fmt.Printf("MySQLAdapter: %+v \n", mySQL)

	result, err := mySQL.DB.Exec(query, r.UserID, r.DataID, r.Version, r.Content)
	if err != nil {
		return errors.E(op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.E(op, err)
	}

	if rowsAffected == 0 {
		return errors.E(op, errors.New("Record was not inserted: 0 rows were affected"))
	}

	return nil
}
