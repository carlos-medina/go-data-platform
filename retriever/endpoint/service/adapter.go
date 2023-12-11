package service

import (
	"database/sql"
	"fmt"

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

func (mySQL MySQLAdapter) GetByDataId(dataId int) (Response, error) {
	const op = errors.Op("gateway.MySQLAdapter.GetByDataId")
	query := fmt.Sprintf("SELECT user_id, data_id, version, content FROM %v WHERE data_id = ?", mySQL.Table)
	var response Response

	row := mySQL.DB.QueryRow(query, dataId)
	if err := row.Scan(&response.UserID, &response.DataID, &response.Version, &response.Content); err != nil {
		if err == sql.ErrNoRows {
			return response, errors.E(op, err)
		}
		return response, errors.E(op, err)
	}

	return response, nil
}
