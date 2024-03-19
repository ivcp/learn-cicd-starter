package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		header      string
		expectError bool
	}{
		{name: "no Authorization header present", header: "", expectError: true},
		{name: "Authorization header melformed", header: "test", expectError: true},
		{name: "good Authorization header", header: "ApiKey test", expectError: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			headers := http.Header{}
			if test.header != "" {
				headers.Add("Authorization", test.header)
			}
			_, err := GetAPIKey(headers)
			if err != nil && !test.expectError {
				t.Errorf("expected no err, but got %q", err)
			}
			if err == nil && test.expectError {
				t.Errorf("expected err, but didn't get one")
			}

			if test.expectError {
				if test.header == "" && !errors.Is(err, ErrNoAuthHeaderIncluded) {
					t.Errorf("expected no header included error, but got: %q", err)
				}
				if test.header != "" && err.Error() != "malformed authorization header" {
					t.Errorf("expected 'malformed authorization header' error, but got: %q", err)
				}
			}
		})
	}
}
