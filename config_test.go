package ddnsman

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_stringDuration_UnmarshalJSON(t *testing.T) {
	tt := []struct {
		name string
		data []byte
		dura time.Duration
		err  string
	}{
		{
			name: "success string",
			data: []byte(`"2m30s"`),
			dura: 2*time.Minute + 30*time.Second,
		},
		{
			name: "success float",
			data: []byte(`150000000000`),
			dura: 2*time.Minute + 30*time.Second,
		},
		{
			name: "error unmarshal",
			data: []byte(`~`),
			err:  "invalid character '~' looking for beginning of value",
		},
		{
			name: "error parse duration",
			data: []byte(`"1h 30m"`),
			err:  `parse time: time: unknown unit "h " in duration "1h 30m"`,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var actual stringDuration
			err := json.Unmarshal(tc.data, &actual)
			assert.Equal(t, tc.dura, time.Duration(actual))
			assertError(t, err, tc.err)
		})
	}
}

func assertError(t *testing.T, theError error, errString string) {
	if errString == "" {
		assert.NoError(t, theError)
		return
	}
	assert.EqualError(t, theError, errString)
}
