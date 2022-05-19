package controller

import (
	"alertCenter/config"
	"alertCenter/controller/common"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DingTalkNotification struct {
	MessageType string                        `json:"msgtype"`
	Text        *DingTalkNotificationText     `json:"text,omitempty"`
	Markdown    *DingTalkNotificationMarkdown `json:"markdown,omitempty"`
	At          *DingTalkNotificationAt       `json:"at,omitempty"`
}

type DingTalkNotificationText struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
type DingTalkNotificationMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type DingTalkNotificationAt struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}
type Query struct {
	Robot string `json:"robot"`
	DingTalkNotification
}
type ReqBody struct {
	Mode  string `json:"mode"`
	Query `json:"query"`
}

// DingTalk 钉钉通道
func DingTalk(c *gin.Context) {
	var (
		url   string
		token string
		dn    ReqBody
	)
	if err := c.ShouldBindJSON(&dn); err != nil {
		common.SendJSON(c, make(map[string]string), 9001, err.Error())
		return
	}

	token = config.DingTalk.Robot[dn.Query.Robot]
	url = "https://oapi.dingtalk.com/robot/send?access_token=" + token
	client := &http.Client{}
	data, err := json.Marshal(&dn.Query.DingTalkNotification)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		common.SendJSON(c, make(map[string]string), 9001, err.Error())
		return
		// panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, _ := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		common.SendJSON(c, make(map[string]string), 9001, err.Error())
		return
		// panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(body))
	common.SendJSON(c, "发送成功.")
}
