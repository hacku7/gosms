/**
* @author: [hk7]
* @Data:   12:27 AM
 */

package aliyun

import (
	"bytes"
	"encoding/json"
	"testing"
)

// AliYunSmsTemplate227130261
// 定义业务需要用到的短信模板
// 我们注册的阿里云短信模板id是227130261
type AliYunSmsTemplate227130261 struct {
	SmsTemplateInterface
}

// GetTemplateId 获取模板ID
func (AliYunSmsTemplate227130261) GetTemplateId() string {
	return "SMS_227130261"
}

// GenerateTemplateParam 根据传入的结构体, 生成短信模板所需的参数
func (AliYunSmsTemplate227130261) GenerateTemplateParam(params interface{}) *string {
	if params == nil {
		return nil
	}

	switch params.(type) {
	case string:
		type templateJsonStruct227130261 struct {
			Code string `json:"code"`
		}
		return converseStructToJsonString(templateJsonStruct227130261{params.(string)})
	}
	return nil
}

// 定义业务需要用到的短信模板
func TestAliYunSMSClientSender_SendSms(t *testing.T) {
	//  初始化短信sender
	const (
		AccessKeyID     = "xxxxxxxxxxxxxxxx"
		AccessKeySecret = "yyyyyyyyyyyyyyyyyyyyyyy"
	)
	credential := AliYunSMSClientCredential{
		AccessKeyId:     AccessKeyID,
		AccessKeySecret: AccessKeySecret,
	}
	sender, err := credential.CreateAliYunSMSClientSender("矩阵", AliYunSmsTemplate227130261{}, &AliYunSMSOptionalConfig{})
	if err != nil {
		t.Error(err.Error())
	}
	// 初始化短信sender

	// 测试发送功能
	resp, err := sender.SendSms([]string{"15555000312"}, nil, nil, "666666")
	if err != nil {
		t.Error(err.Error())
	}
	//println(resp.IsRequestSuccess())
	//println(resp.GetBizId())
	//println(resp.IsBizIdValid())
	println(resp.GetRequestStatusCode())
	//println(resp.GetRequestStatusRequestId())
	//println(*resp.QueryRequestStatusErrorInfo())
	//  测试发送功能
}

func converseStructToJsonString(v interface{}) *string {
	jData, err := json.Marshal(v)
	if err != nil {
		return nil
	}

	var out bytes.Buffer
	err = json.Indent(&out, jData, "", "\t")
	if err != nil {
		return nil
	}

	res := out.String()
	return &res
}
