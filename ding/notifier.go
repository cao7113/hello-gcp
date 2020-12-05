package ding

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Notifying interface {
	Send(title, msg string, mobiles []string)
}

type Notifier struct {
}

func (n *Notifier) Send(title, msg string, mobiles []string) {
	content := fmt.Sprintf("%s\n%s", title, msg)
	err := SendText(content, false, mobiles)
	if err != nil {
		logrus.Error("Notifier Send, error: ", err.Error())
	}
}
