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

func TestAppAction(t *testing.T) {
	cases := []struct {
		name        string
		args        []string
		expectedErr bool
	}{
		{
			name:        "Valid timestamp",
			args:        []string{"tsd", "1833561154"},
			expectedErr: false,
		},
		{
			name:        "Invalid timestamp",
			args:        []string{"tsd", "asdf"},
			expectedErr: true,
		},
		{
			name:        "No arguments",
			args:        []string{"tsd"},
			expectedErr: true,
		},
		{
			name:        "With flag",
			args:        []string{"tsd", "-t", "1833561154"},
			expectedErr: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			app := NewApp()
			err := app.Run(c.args)
			if c.expectedErr {
				require.Error(t, err, "Error case")
			} else {
				require.NoError(t, err, "No error case")
			}
		})
	}
}
