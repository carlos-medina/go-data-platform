package endpoint

import (
	"testing"

	"github.com/arquivei/foundationkit/errors"
	"github.com/stretchr/testify/assert"
)

func TestDecodeInput(t *testing.T) {
	tests := []struct {
		name          string
		input         []byte
		expecteResult Record
		expectedError string
	}{
		{
			name:          "Failure: Simple string",
			input:         []byte("aldsls"),
			expectedError: "invalid character 'a' looking for beginning of value",
		},
		{
			name:          "Failure - Missing UserID",
			input:         []byte("{\"data_id\":1,\"version\":1,\"content\":\"soap\"}"),
			expectedError: "Zero value fields: [UserID]",
		},
		{
			name:          "Failure - Missing DataID",
			input:         []byte("{\"user_id\":1,\"version\":1,\"content\":\"soap\"}"),
			expectedError: "Zero value fields: [DataID]",
		},
		{
			name:          "Failure - Missing Version",
			input:         []byte("{\"user_id\":1,\"data_id\":1,\"content\":\"soap\"}"),
			expectedError: "Zero value fields: [Version]",
		},
		{
			name:          "Failure - Missing Content",
			input:         []byte("{\"user_id\":1,\"data_id\":1,\"version\":1}"),
			expectedError: "Zero value fields: [Content]",
		},
		{
			name:          "Failure - Missing DataID and Content",
			input:         []byte("{\"user_id\":1,\"version\":1}"),
			expectedError: "Zero value fields: [DataID Content]",
		},
		{
			name:          "Failure - Missing all fields",
			input:         []byte("{}"),
			expectedError: "Zero value fields: [UserID DataID Version Content]",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			record, err := DecodeInput(test.input)
			if test.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, test.expecteResult, record)
			} else {
				assert.EqualError(t, errors.GetRootErrorWithKV(err), test.expectedError)
			}
		})
	}
}

func TestUnmarshalInputData(t *testing.T) {
	tests := []struct {
		name          string
		input         []byte
		expecteResult Record
		expectedError string
	}{
		{
			name:          "Failure: Simple string",
			input:         []byte("aldsls"),
			expectedError: "invalid character 'a' looking for beginning of value",
		},
		{
			name:          "Failure: Empty string",
			input:         []byte(""),
			expectedError: "unexpected end of JSON input",
		},
		{
			name:          "Failure: Incorrect type int when expected string",
			input:         []byte("{\"user_id\":1,\"data_id\":1,\"version\":1,\"content\":1}"),
			expectedError: "json: cannot unmarshal number into Go struct field Record.content of type string",
		},
		{
			name:  "Success - Empty map",
			input: []byte("{}"),
			expecteResult: Record{
				UserID:  0,
				DataID:  0,
				Version: 0,
				Content: "",
			},
		},
		{
			name:  "Success - Missing value from key",
			input: []byte("{\"user_id\":1,\"data_id\":1,\"version\":1}"),
			expecteResult: Record{
				UserID:  1,
				DataID:  1,
				Version: 1,
				Content: "",
			},
		},
		{
			name:  "Success - All keys have values",
			input: []byte("{\"user_id\":1,\"data_id\":1,\"version\":1,\"content\":\"soap\"}"),
			expecteResult: Record{
				UserID:  1,
				DataID:  1,
				Version: 1,
				Content: "soap",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			record, err := unmarshalInput(test.input)
			if test.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, test.expecteResult, record)
			} else {
				assert.EqualError(t, errors.GetRootErrorWithKV(err), test.expectedError)
			}
		})
	}
}
