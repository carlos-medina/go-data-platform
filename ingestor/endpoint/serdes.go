package endpoint

import (
	"encoding/json"
	"fmt"

	"github.com/arquivei/foundationkit/errors"
)

type Record struct {
	UserID  int    `json:"user_id"`
	DataID  int    `json:"data_id"`
	Version int    `json:"version"`
	Content string `json:"content"`
}

// DecodeInput validates if all Record fields have non zero values
// after doing the input unmarshal
func DecodeInput(input []byte) (Record, error) {
	const op = errors.Op("endpoint.DecodeInput")

	record, err := unmarshalInput(input)

	if err != nil {
		return Record{}, errors.E(op, err)
	}

	var zeroValueFileds []string

	if record.UserID == 0 {
		zeroValueFileds = append(zeroValueFileds, "UserID")
	}
	if record.DataID == 0 {
		zeroValueFileds = append(zeroValueFileds, "DataID")
	}
	if record.Version == 0 {
		zeroValueFileds = append(zeroValueFileds, "Version")
	}
	if record.Content == "" {
		zeroValueFileds = append(zeroValueFileds, "Content")
	}

	if zeroValueFileds != nil {
		return Record{}, errors.E(op, errors.New(fmt.Sprintf("Zero value fields: %v", zeroValueFileds)))
	}

	return record, nil
}

// unmarshalInput unmarshals the input in the record struct
func unmarshalInput(input []byte) (Record, error) {
	const op = errors.Op("endpoint.UnmarshalInputData")

	var record Record

	err := json.Unmarshal(input, &record)
	if err != nil {
		return Record{}, errors.E(op, err)
	}

	return record, nil
}
