package converter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimestampToDate(t *testing.T) {
	cases := []struct {
		name        string
		value       string
		expected    time.Time
		expectedErr bool
	}{
		{
			name:        "valid timestamp",
			value:       "1733561154",
			expected:    time.Date(2024, 12, 7, 8, 45, 54, 0, time.UTC),
			expectedErr: false,
		},
		{
			name:        "invalid timestamp",
			value:       "1733561154a",
			expected:    time.Time{},
			expectedErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(*testing.T) {
			actual, err := TimestampToDate(c.value)
			if c.expectedErr {
				assert.Equal(t, actual, c.expected)
				require.Error(t, err, "")
			} else {

				assert.Equal(t, actual, c.expected)
				require.NoError(t, err)
			}
		})
	}
}
