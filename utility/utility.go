package utility

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

func GetToken() string {
	client := resty.New()
	payload := map[string]interface{}{
		"email":    "zjf15821145685@163.com",
		"password": "1234",
	}
	resp, _ := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("http://47.96.125.121:8080/user/login")
	var result map[string]string
	_ = json.Unmarshal(resp.Body(), &result)
	token := result["data"]
	return token
}
