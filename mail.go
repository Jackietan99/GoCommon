package common

import (
	"net/smtp"
	"strings"
)

/**
* 发送系统邮件
* @host 邮件服务器地址
* @user 登录账户
* @psw 登录密码
* @mailuser 指定发送者邮箱
* @mailType string 发送邮件的格式类型：decault纯文本 text； 网页 html,
* @tomail 接收者邮箱地址，多个用英文分号分开隔开
* @subject 发送邮件的标题
* @content 发送邮件的内容
* @return int 发送状态 200成功,其他状态表示失败
   string 发送结束后的相关信息
*/

func SendMail(host, user, psw, mailuser, mailType string, tomail string, subject string, content string) (int, string) {
	send_status := 100
	send_msg := "邮件发送失败"

	head := ""
	foot := ""
	//如果是网页格式的，增加html头
	if mailType == "html" {
		head = "<!DOCTYPE html>"
		head = head + "<html>"
		head = head + "<head>"
		head = head + "<meta charset=\"UTF-8\" />"
		head = head + "</head>"
		head = head + "<body>"
		foot = "</body>"
		foot = foot + "</html>"
	}
	content = head + content + foot

	//查看是否有端口
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, psw, hp[0])
	send_to := strings.Split(tomail, ";")
	err := smtp.SendMail(host, auth, mailuser, send_to, []byte("Subject: "+subject+"\r\n"+content))
	if err == nil {
		send_status = 200
		send_msg = "邮件发送成功"
	} else {
		send_msg = err.Error()
	}
	return send_status, send_msg
}
