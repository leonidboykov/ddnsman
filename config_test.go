package ddnsman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfiguration_shoutrrrURLs(t *testing.T) {
	tt := []struct {
		name  string
		addrs []ShoutrrrNotifications
		urls  []string
		err   string
	}{
		{
			name: "success",
			addrs: []ShoutrrrNotifications{
				{
					URL: "https://example.com",
					Settings: map[string]string{
						"key1": "val1",
						"key2": "val2",
					},
				},
				{
					URL: "https://example.org",
					Settings: map[string]string{
						"key3": "val3",
						"key4": "val4",
					},
				},
			},
			urls: []string{
				"https://example.com?key1=val1&key2=val2",
				"https://example.org?key3=val3&key4=val4",
			},
		},
		{
			name: "error",
			addrs: []ShoutrrrNotifications{
				{URL: ":~"},
			},
			err: `parse url: parse ":~": missing protocol scheme`,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := Configuration{ShoutrrrAddrs: tc.addrs}
			urls, err := c.shoutrrrURLs()
			assert.Equal(t, tc.urls, urls)
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
