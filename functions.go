// Contains an HTTP Cloud Function.
package hellogcp

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/cao7113/hellogcp/ding"
	"github.com/cao7113/hellogcp/ping"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("Error loading .env file")
	}
}

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
			logrus.Infof("json.NewDecoder: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	switch d.Message {
	case "":
		fmt.Fprint(w, "Hello, Blank!")
		return
	case "fatal":
		logrus.Fatal("hit fatal")
		return
	case "panic":
		panic("hit panic")
		return
	}

	logrus.Infof("received message: %s at %s", d.Message, time.Now().Format(time.RFC3339Nano))
	fmt.Fprintf(w, "Hello %s!", html.EscapeString(d.Message))
}

func Ding(w http.ResponseWriter, r *http.Request) {
	// ding at share-up message
	n := &ding.Notifier{}
	title := fmt.Sprintf("Up %s", "now")
	content := fmt.Sprintf("received at %s \n method: %s url: %s agent: %s", time.Now().Format(time.RFC3339), r.Method, r.URL.String(), r.UserAgent())
	n.Send(title, content, []string{})
	fmt.Fprintf(w, "test ding done")
}
