package cloudfunc

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestPint(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/ping", nil)
	w := httptest.NewRecorder()
	Ping(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}
