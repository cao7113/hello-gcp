package hellogcp

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

func TestHelloHTTPSystem(t *testing.T) {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		t.SkipNow()
	}
	urlString := baseURL
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	testURL, err := url.Parse(urlString)
	if err != nil {
		t.Fatalf("url.Parse(%q): %v", urlString, err)
	}

	tests := []struct {
		body string
		want string
	}{
		{body: `{"message": ""}`, want: "Hello, Blank!"},
		{body: `{"message": "Gopher"}`, want: "Hello Gopher!"},
	}

	for _, test := range tests {
		req := &http.Request{
			Method: http.MethodPost,
			Body:   ioutil.NopCloser(strings.NewReader(test.body)),
			URL:    testURL,
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("HelloHTTP http.Get: %v", err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("HelloHTTP ioutil.ReadAll: %v", err)
		}
		if got := string(body); got != test.want {
			t.Errorf("HelloHTTP(%q) = %q, want %q", test.body, got, test.want)
		}
	}
}
