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

	"github.com/cao7113/hellogcp/ping"
	"github.com/sirupsen/logrus"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	ping.Ping(w, r)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json:"message"`
	}

	//w.Header().Add("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		switch err {
		case io.EOF:
			fmt.Fprint(w, "Hello, EOF!")
			return
		default:
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	switch d.Message {
	case "":
		fmt.Fprint(w, "Hello, Blank!")
		return
	case "fatal":
		log.Fatal("hit fatal")
		return
	case "panic":
		panic("hit panic")
		return
	}

	logrus.Infof("received message: %s at %s", d.Message, time.Now().Format(time.RFC3339Nano))
	fmt.Fprintf(w, "Hello %s!", html.EscapeString(d.Message))
}
