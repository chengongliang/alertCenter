package controller

import (
	"alertCenter/config"
	"alertCenter/controller/common"
	"alertCenter/model"
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
)

// SMS 短信通道
func SMS(c *gin.Context) {
	var (
		err     error
		reqBody model.ReqBody
		mobile  string
		content string
	)
	if err = c.ShouldBindJSON(&reqBody); err != nil {
		common.SendJSON(c, make(map[string]string, 0), 9001, err.Error())
		return
	}
	mobile = reqBody.Query["mobile"].(string)
	content = reqBody.Query["content"].(string)
	client, err := dysmsapi.NewClientWithAccessKey(config.AliSMS.RegionID, config.AliSMS.AccessKeyID, config.AliSMS.AccessSecret)
	m := map[string]string{"msg": content}
	mjson, _ := json.Marshal(m)
	mString := string(mjson)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = mobile
	request.SignName = config.AliSMS.SignName
	request.TemplateCode = config.AliSMS.TemplateCode
	request.TemplateParam = mString

	response, err := client.SendSms(request)
	if err != nil {
		common.SendJSON(c, make(map[string]string, 0), 9001, err.Error())
		return
	}
	common.SendJSON(c, response)
}
