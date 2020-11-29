package hellogcp

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"testing"
)

func TestGo(t *testing.T) {
	logrus.Infof("version: %s", runtime.Version())
}
