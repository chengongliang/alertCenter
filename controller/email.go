package controller

import (
	"alertCenter/config"
	"alertCenter/controller/common"
	"alertCenter/model"
	"crypto/tls"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"

	"github.com/gin-gonic/gin"
)

// Email 邮件通道
func Email(c *gin.Context) {
	var (
		to       string
		subject  string
		mailType string
		content  string
		sender   string
		toList   []string
		err      error
		reqBody  model.ReqBody
	)

	if err = c.ShouldBindJSON(&reqBody); err != nil {
		common.SendJSON(c, make(map[string]string), 9001, err.Error())
		return
	}
	sender = reqBody.Query["sender"].(string)
	to = reqBody.Query["to"].(string)
	subject = reqBody.Query["subject"].(string)
	content = reqBody.Query["content"].(string)
	mailType = "text/plain"
	if reqBody.Query["mail_type"].(string) == "html" {
		mailType = "text/html"
	}
	if s, ok := config.Email[sender]; ok {
		port := 587
		if port, err = strconv.Atoi(s.SMTPPort); err != nil {
			common.SendJSON(c, make(map[string]string), 9001, err.Error())
		}
		toList = strings.Split(strings.TrimSuffix(to, ","), ",")
		m := gomail.NewMessage()
		m.SetHeader("From", s.From)
		m.SetHeader("To", toList...)
		m.SetHeader("Subject", subject)
		m.SetBody(mailType, content)
		d := gomail.NewDialer(s.SMTPServer, port, s.UserName, s.Password)
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		fmt.Println(d)
		if err := d.DialAndSend(m); err != nil {
			fmt.Println(err)
			d.DialAndSend(m)
		}
		common.SendJSON(c, "邮件发送成功")
	} else {
		keys := make([]string, 0, len(config.Email))
		for k := range config.Email {
			keys = append(keys, k)
		}
		res := fmt.Sprintf("%v 未配置！已配置sender: %v", sender, keys)
		common.SendJSON(c, make(map[string]string), 9001, res)
	}
}
