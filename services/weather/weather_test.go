package weather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDescribeTemp(t *testing.T) {
	var myTests = map[string]struct {
		temp         float64
		expectedDesc string
	}{
		"Given a temp higher than 80 expect to return 'hot'": {
			temp:         100,
			expectedDesc: "hot",
		},
		"Given a temp between 50 and 80 expect to return 'moderate'": {
			temp:         65,
			expectedDesc: "moderate",
		},
		"Given a temp below 50 expect to return 'cold'": {
			temp:         25,
			expectedDesc: "cold",
		},
	}
	for _, tc := range myTests {
		assert.Equal(t, tc.expectedDesc, describeTemp(tc.temp))
	}
}
