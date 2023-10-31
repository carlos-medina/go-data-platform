package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStr(t *testing.T) {
	tests := []struct {
		name           string
		seed           int64
		lenght         int
		expectedResult string
	}{
		{
			name:           "seed = 0; lenght = 10",
			seed:           0,
			lenght:         10,
			expectedResult: "YODGPVRCND",
		},
		{
			name:           "lenght = 0",
			lenght:         0,
			expectedResult: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RandStr(test.seed, test.lenght)
			assert.Equal(t, test.expectedResult, result)
		})
	}
}
