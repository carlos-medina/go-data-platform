package endpoint

import (
	"testing"

	"github.com/arquivei/foundationkit/errors"
	"github.com/stretchr/testify/assert"
)

func TestDecodeInput(t *testing.T) {
	tests := []struct {
		name          string
		inputBytes    []byte
		expecteResult Input
		expectedError string
	}{
		{
			name:          "Failure: Simple string",
			inputBytes:    []byte("aldsls"),
			expectedError: "invalid character 'a' looking for beginning of value",
		},
		{
			name:          "Failure - Missing UserID",
			inputBytes:    []byte("{\"data_id\":1,\"version\":1,\"content\":\"soap\"}"),
			expectedError: "Zero value fields: [UserID]",
		},
		{
			name:          "Failure - Missing DataID",
			inputBytes:    []byte("{\"user_id\":1,\"version\":1,\"content\":\"soap\"}"),
			expectedError: "Zero value fields: [DataID]",
		},
		{
			name:          "Failure - Missing Version",
			inputBytes:    []byte("{\"user_id\":1,\"data_id\":1,\"content\":\"soap\"}"),
			expectedError: "Zero value fields: [Version]",
		},
		{
			name:          "Failure - Missing Content",
			inputBytes:    []byte("{\"user_id\":1,\"data_id\":1,\"version\":1}"),
			expectedError: "Zero value fields: [Content]",
		},
		{
			name:          "Failure - Missing DataID and Content",
			inputBytes:    []byte("{\"user_id\":1,\"version\":1}"),
			expectedError: "Zero value fields: [DataID Content]",
		},
		{
			name:          "Failure - Missing all fields",
			inputBytes:    []byte("{}"),
			expectedError: "Zero value fields: [UserID DataID Version Content]",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inputData, err := DecodeInput(test.inputBytes)
			if test.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, test.expecteResult, inputData)
			} else {
				assert.EqualError(t, errors.GetRootErrorWithKV(err), test.expectedError)
			}
		})
	}
}

func TestUnmarshalInputData(t *testing.T) {
	tests := []struct {
		name          string
		inputBytes    []byte
		expecteResult Input
		expectedError string
	}{
		{
			name:          "Failure: Simple string",
			inputBytes:    []byte("aldsls"),
			expectedError: "invalid character 'a' looking for beginning of value",
		},
		{
			name:          "Failure: Empty string",
			inputBytes:    []byte(""),
			expectedError: "unexpected end of JSON input",
		},
		{
			name:          "Failure: Incorrect type int when expected string",
			inputBytes:    []byte("{\"user_id\":1,\"data_id\":1,\"version\":1,\"content\":1}"),
			expectedError: "json: cannot unmarshal number into Go struct field Input.content of type string",
		},
		{
			name:       "Success - Empty map",
			inputBytes: []byte("{}"),
			expecteResult: Input{
				UserID:  0,
				DataID:  0,
				Version: 0,
				Content: "",
			},
		},
		{
			name:       "Success - Missing value from key",
			inputBytes: []byte("{\"user_id\":1,\"data_id\":1,\"version\":1}"),
			expecteResult: Input{
				UserID:  1,
				DataID:  1,
				Version: 1,
				Content: "",
			},
		},
		{
			name:       "Success - All keys have values",
			inputBytes: []byte("{\"user_id\":1,\"data_id\":1,\"version\":1,\"content\":\"soap\"}"),
			expecteResult: Input{
				UserID:  1,
				DataID:  1,
				Version: 1,
				Content: "soap",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inputData, err := unmarshalInput(test.inputBytes)
			if test.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, test.expecteResult, inputData)
			} else {
				assert.EqualError(t, errors.GetRootErrorWithKV(err), test.expectedError)
			}
		})
	}
}
