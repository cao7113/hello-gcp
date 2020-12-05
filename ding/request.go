package ding

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

// http-url?access_token=xx-token
const DingURLEnvKey = "DING_TALK_URL_WITH_TOKEN"

func getDingURL() string {
	return os.Getenv(DingURLEnvKey)
}

func sendRequest(msg []byte) error {
	aurl := getDingURL()
	if aurl == "" {
		logrus.Warnf("skip ding message: %s", string(msg))
		return nil
	}
	r := bytes.NewReader(msg)
	req, err := http.NewRequest("POST", aurl, r)
	if err != nil {
		logrus.Error("DingTalk new request error:", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error("DingTalk do request error:", err)
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	logrus.Info("DingTalk response body:", string(body))
	return nil
}
