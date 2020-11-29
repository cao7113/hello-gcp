package hellogcp

import (
	"github.com/sirupsen/logrus"
	"time"
)

func ModLog(){
	logrus.Infof("log at %s", time.Now().Format(time.RFC3339Nano))
}
