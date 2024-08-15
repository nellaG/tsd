package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewApp(t *testing.T) {
	cases := []struct {
		name          string
		expectedName  string
		expectedFlags int
	}{
		{
			name:          "Default app configuration",
			expectedName:  "tsd",
			expectedFlags: 1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			app := NewApp()

			require.Equal(t, app.Name, c.expectedName, "App name should match")
			require.Len(t, app.Flags, c.expectedFlags, "# of flags should match")
			require.NotNil(t, app.Action, "Action should be set")
		})
	}

}
