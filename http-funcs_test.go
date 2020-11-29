package hellogcp

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/ping", nil)
	//req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	Hello(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Header.Get("Content-Type"))
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Hello, EOF!", string(body))
}
