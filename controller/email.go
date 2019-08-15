package controller

import (
	"alertCenter/config"
	"alertCenter/controller/common"
	"alertCenter/model"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/gin-gonic/gin"
)

// Email 邮件通道
func Email(c *gin.Context) {
	var (
		to       string
		msg      string
		addr     string
		mailFrom string
		subject  string
		mailType string
		content  string
		auth     smtp.Auth
		toList   []string
		err      error
		reqBody  model.ReqBody
	)

	if err = c.ShouldBindJSON(&reqBody); err != nil {
		common.SendJSON(c, make(map[string]string, 0), 9001, err.Error())
		return
	}
	to = reqBody.Query["to"].(string)
	subject = reqBody.Query["subject"].(string)
	content = reqBody.Query["content"].(string)
	mailType = "text/plain"
	if reqBody.Query["mail_type"].(string) == "html" {
		mailType = "text/html"
	}
	mailFrom = config.Email.From
	addr = config.Email.SMTPServer + ":" + config.Email.SMTPPort
	auth = smtp.PlainAuth("", config.Email.UserName, config.Email.Password, config.Email.SMTPServer)
	msg = fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\nContent-Type: %s; charset=UTF-8\r\n\r\n%s", to, mailFrom, subject, mailType, content)
	toList = strings.Split(strings.TrimSuffix(to, ","), ",")
	err = smtp.SendMail(
		addr,
		auth,
		config.Email.UserName,
		toList,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err.Error())
		for _, receiver := range toList {
			ToEach := []string{receiver}
			smtp.SendMail(
				addr,
				auth,
				config.Email.UserName,
				ToEach,
				[]byte(msg),
			)
		}
	}
	common.SendJSON(c, "邮件发送成功")
}
