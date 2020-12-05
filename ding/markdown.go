package ding

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Markdown struct {
	MsgType  string      `json:"msgtype"`
	Markdown MarkdownMsg `json:"markdown"`
	At       At          `json:"at"`
}

type MarkdownMsg struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// SendMarkdown: 在text内容里要有@手机号, mobiles才奏效
func SendMarkdown(title, text string, isAtAll bool, mobiles []string) error {
	if len(mobiles) > 0 {
		mobilesText := ""
		for _, m := range mobiles {
			mobilesText = fmt.Sprintf("%s@%s", mobilesText, m)
		}
		text = fmt.Sprintf("%s %s", text, mobilesText)
	}
	logrus.Infof("DingTalk Markdown text: %s", text)

	msg := Markdown{
		MsgType: "markdown",
		Markdown: MarkdownMsg{
			Title: title,
			Text:  text,
		},
		At: At{
			AtMobiles: mobiles,
			IsAtAll:   isAtAll,
		},
	}
	msgValue, err := json.Marshal(msg)
	if err != nil {
		logrus.Errorf("DingTalk Markdown json marshal error: %s", err.Error())
		return err
	}

	return sendRequest(msgValue)
}
