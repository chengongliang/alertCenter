> 集中告警通道

- 阿里短信通道
- 钉钉报警
- 邮件报警

> 用法

- 钉钉
  ```
  curl -X POST \
  http://localhost:9001/dingtalk \
  -H 'token: xxx' \
  -d '{
    "mode": "",
    "query": {
    	"robot": "op",
    	"content": "xxx"
    }
  }'
  ```

- 邮件
  ```
  curl -X POST \
  http://10.7.103.152:9001/mail \
  -H 'token: xxx' \
  -d '{
    "mode": "",
    "query": {
    	"mail_type": "text", # [text|html]
    	"subject":"告警通知",
    	"content": "xxx",
    	"to":"a@xx.com,b@xx.com" # 多个邮箱用`,`分割
    }
  }'
  ```

- 短信
  ```
  curl -X POST \
  http://localhost:9001/sms \
  -H 'token: HQZXLTAIL9AyaIgBrtWs' \
  -d '{
    "mode": "",
    "query": {
    	"mobile": "123,456", # 多个号码用`,`分割
    	"content": "xxx"
    }
  }'
  ```

> TODO

- 数据库记录