package message

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"software/common"
)

var client *dysmsapi.Client

type smsMessage struct {
	Code string `json:"code"`
}

func BackendAuth(keyId string, keySecret string) (err error) {
	client, err = dysmsapi.NewClientWithAccessKey("cn-hangzhou", keyId, keySecret)
	return err
}

func SendSms(cellphone string, bindcode string) error {
	if client == nil {
		return common.SmsErr
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = cellphone
	request.SignName = "虚拟车辆管理系统"
	request.TemplateCode = "SMS_191190567"

	message := smsMessage{Code: bindcode}
	template, err := json.Marshal(message)
	request.TemplateParam = string(template)
	_, err = client.SendSms(request)
	if err != nil {
		return err
	}
	return nil
}
