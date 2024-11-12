package auth_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestAuth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	req.Header.Set("Authorization", "Test-Header")

	type WantType struct {
		text string
		err  string
	}

	tests := map[string]struct {
		req    http.Request
		header string
		want   WantType
	}{
		"no-header":    {header: "", req: *req, want: WantType{text: "", err: "no authorization header included"}},
		"wrong-header": {header: "hello", req: *req, want: WantType{text: "", err: "malformed authorization header"}},
		"success":      {header: "hello", req: *req, want: WantType{text: "", err: "duupa"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.req.Header.Set("Authorization", tc.header)
			val, err := auth.GetAPIKey(tc.req.Header)
			fmt.Printf("error: %#v", err)

			if val != tc.want.text {
				t.Fatalf("expected: %#v, got: %#v", tc.want.text, val)
			}

			if !strings.Contains(fmt.Sprint(err), tc.want.err) {
				t.Fatalf("expected: %#v, got: %#v", tc.want.err, err)
			}
		})
	}
}
