package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		want          string
		wantErrString string
	}{
		{
			name:    "valid api key",
			headers: http.Header{"Authorization": []string{"ApiKey 12345"}},
			want:    "12345",
		},
		{
			name:          "no auth header",
			headers:       http.Header{},
			wantErrString: "no authorization header included",
		},
		{
			name:          "malformed auth header (no ApiKey prefix)",
			headers:       http.Header{"Authorization": []string{"Bearer 12345"}},
			wantErrString: "malformed authorization header",
		},
		{
			name:          "malformed auth header (too short)",
			headers:       http.Header{"Authorization": []string{"ApiKey"}},
			wantErrString: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) && err.Error() != tt.wantErrString {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErrString)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
