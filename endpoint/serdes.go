package endpoint

import (
	"encoding/json"
	"fmt"

	"github.com/arquivei/foundationkit/errors"
)

type Input struct {
	UserID  int    `json:"user_id"`
	DataID  int    `json:"data_id"`
	Version int    `json:"version"`
	Content string `json:"content"`
}

// DecodeInput validates if all Input fields have non zero values
// after doing the data unmarshal
func DecodeInput(inputBytes []byte) (Input, error) {
	const op = errors.Op("endpoint.DecodeInput")

	input, err := unmarshalInput(inputBytes)

	if err != nil {
		return Input{}, errors.E(op, err)
	}

	var zeroValueFileds []string

	if input.UserID == 0 {
		zeroValueFileds = append(zeroValueFileds, "UserID")
	}
	if input.DataID == 0 {
		zeroValueFileds = append(zeroValueFileds, "DataID")
	}
	if input.Version == 0 {
		zeroValueFileds = append(zeroValueFileds, "Version")
	}
	if input.Content == "" {
		zeroValueFileds = append(zeroValueFileds, "Content")
	}

	if zeroValueFileds != nil {
		return Input{}, errors.E(op, errors.New(fmt.Sprintf("Zero value fields: %v", zeroValueFileds)))
	}

	return input, nil
}

// unmarshalInput unmarshal the input in its struct
func unmarshalInput(inputBytes []byte) (Input, error) {
	const op = errors.Op("endpoint.UnmarshalInputData")

	var inputData Input

	err := json.Unmarshal(inputBytes, &inputData)
	if err != nil {
		return Input{}, errors.E(op, err)
	}

	return inputData, nil
}
