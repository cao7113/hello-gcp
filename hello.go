// Contains an HTTP Cloud Function.
package hellogcp

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"time"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		switch err {
		case io.EOF:
			fmt.Fprint(w, "Hello World!")
			return
		default:
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	switch d.Message {
	case "":
		fmt.Fprint(w, "Hello World!")
		return
	case "fatal":
		log.Fatal("hit fatal")
		return
	case "panic":
		panic("hit panic")
		return
	}

	log.Printf("received message: %s at %s", d.Message, time.Now().Format(time.RFC3339Nano))
	fmt.Fprint(w, html.EscapeString(d.Message))
}
