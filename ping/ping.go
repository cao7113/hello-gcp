// Contains an HTTP Cloud Function.
package ping

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// Ping ping-pong test
func Ping(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("%s request %s with user-agent: %s", r.Method, r.URL, r.UserAgent())
	fmt.Fprintf(w,"pong at %s", time.Now().Format(time.RFC3339Nano))
}
