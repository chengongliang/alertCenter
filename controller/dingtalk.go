package controller

import (
	"alertCenter/config"
	"alertCenter/controller/common"
	"alertCenter/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"strings"
)

// DingTalk 钉钉通道
func DingTalk(c *gin.Context) {
	var (
		url       string
		token     string
		data      string
		content   string
		robotName string
		reqBody   model.ReqBody
	)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		common.SendJSON(c, make(map[string]string, 0), 9001, err.Error())
		return
	}
	robotName = reqBody.Query["robot"].(string)
	content = reqBody.Query["content"].(string)
	token = config.DingTalk.Robot[robotName]
	url = "https://oapi.dingtalk.com/robot/send?access_token=" + token
	data = fmt.Sprintf(`{"msgtype": "text", "text": {"content": %s}}`, content)
	fmt.Println(url, data)
	os.Exit(1)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
