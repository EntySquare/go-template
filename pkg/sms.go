package pkg

import (
	"errors"
	"fmt"
	"github.com/vonage/vonage-go-sdk"
	client "github.com/yunpian/yunpian-go-sdk/sdk"
	"go-template/config"
	"go-template/routing/types"
	"gopkg.in/gomail.v2"
)

// 发送国际短信
func SendVonageCodeMsg(msgParam types.SengMsg) error {
	auth := vonage.CreateAuthFromKeySecret("0f32143c", "LF9C0YzY1kgXHVQY")
	smsClient := vonage.NewSMSClient(auth)
	opts := vonage.SMSOpts{
		Type: "unicode",
	}
	response, errResp, err := smsClient.Send("[ISPAY]", msgParam.Area+msgParam.Phone, "Verification code is: "+msgParam.Msg, opts)
	if err != nil {
		return err
	}
	fmt.Println(response)
	fmt.Println(errResp)
	fmt.Println(err)
	if response.Messages[0].Status == "0" {
		fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
	} else {
		fmt.Println("SendMessage Error code " + errResp.Messages[0].Status + ": " + errResp.Messages[0].ErrorText)
		return errors.New(errResp.Messages[0].ErrorText)
	}
	return nil
}

// 发送国内短信
func SendCheckCodeMessage(msgParam types.SengMsg) error {
	// 发送短信
	ypClient := client.New("ec557d72a53ef29f0aa0c39e79d59814")
	param := client.NewParam(2)
	param[client.MOBILE] = msgParam.Phone
	param[client.TEXT] = "【 TreasureBox】您的手机验证码是" + msgParam.Msg + "。本条信息无需回复"
	r := ypClient.Sms().SingleSend(param)
	if r.Code != 0 {
		return errors.New(r.Msg)
	}
	return nil
}

// 发送邮件
func SendEmailMessage(toEmail string) error {
	validCode := RandomNumber(6)
	// 发送电子邮件
	fromEmail := config.FromEmail
	password := config.EmailPassword
	mail := gomail.NewMessage()
	mail.SetHeader("From", fromEmail)
	mail.SetHeader("To", toEmail)
	mail.SetHeader("Subject", "新注册邮箱验证码 New registration verification code")
	text := "Please find the verification code: <strong><span style=\"font-size: 18px;\">" + validCode + "</span></strong>. Valid in 3 minutes"
	mail.SetBody("text/html", text)
	d := gomail.NewDialer("smtp.gmail.com", 465, fromEmail, password) //SMTP： 使用 Gmail 来发送电子邮件
	if err := d.DialAndSend(mail); err != nil {
		fmt.Println("Failed to send email:", err)
		return err
	}
	return nil
}
