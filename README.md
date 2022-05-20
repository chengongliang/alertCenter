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
    	"msgtype": "markdown", # [text|markdown]
      # "text": {  # msgtype: text
      #     "content":"我就是我, @XXX 是不一样的烟火"
      # },
      "markdown": {
         "title":"杭州天气",
         "text": "#### 杭州天气 @150XXXXXXXX \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n"
      },
        "at": {
            "atMobiles": [
                "150XXXXXXXX"
            ],
            "atUserIds": [
                "user123"
            ],
            "isAtAll": false
        }
      }
    }'
    ```

- 邮件
  ```
  curl -X POST \
  http://localhost:9001/mail \
  -H 'token: xxx' \
  -d '{
    "mode": "",
    "query": {
      "sender": "sender1",
    	"msgtype": "text", # [text|html]
    	"subject":"告警通知",
      "sender": "sender1",
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
