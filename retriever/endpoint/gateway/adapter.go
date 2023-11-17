package gateway

import (
	"database/sql"
	"fmt"

	"github.com/carlos-medina/go-data-platform/retriever/endpoint"

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

func (mySQL MySQLAdapter) GetByDataId(dataId int) (endpoint.Record, error) {
	const op = errors.Op("gateway.MySQLAdapter.GetByDataId")
	query := fmt.Sprintf("SELECT user_id, data_id, version, content FROM %v WHERE data_id = ?", mySQL.Table)
	var record endpoint.Record

	row := mySQL.DB.QueryRow(query, dataId)
	if err := row.Scan(&record.UserID, &record.DataID, &record.Version, &record.Content); err != nil {
		if err == sql.ErrNoRows {
			return record, errors.E(op, err)
		}
		return record, errors.E(op, err)
	}

	return record, nil
}
