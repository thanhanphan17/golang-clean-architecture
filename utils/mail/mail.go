package utils

import (
	"bytes"
	"errors"
	"fmt"
	"go-clean-architecture/config"
	"html/template"
	"net/smtp"
	"strings"
)

type MailConfig struct {
	MailFrom   string
	MailServer string
	MailPort   int
	MailPass   string
}

func Init() MailConfig {
	appConfig := config.InitLoadAppConf()

	return MailConfig{
		MailFrom:   appConfig.MailFrom,
		MailServer: appConfig.MailServer,
		MailPort:   appConfig.MailPort,
		MailPass:   appConfig.MailPass,
	}
}

func Send(to, subject, templatePath string, data interface{}) error {
	cfg := Init()

	if len(to) == 0 || len(subject) == 0 || len(templatePath) == 0 {
		return errors.New("to, subject, templatePath can not empty")
	}

	body, err := parseTemplate(templatePath, data)
	if err != nil {
		return err
	}

	var msgs []string

	msgs = append(msgs, "From: Microlab<"+cfg.MailFrom+">\r")
	msgs = append(msgs, "To: "+to+"\r")
	msgs = append(msgs, "Subject: "+subject+"\r")
	msgs = append(msgs, "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n")
	msgs = append(msgs, body+"\r")

	msg := []byte(strings.Join(msgs, "\n"))
	mailAuth := fmt.Sprintf("%s:%d", cfg.MailServer, cfg.MailPort) //465 - 587

	err = smtp.SendMail(mailAuth,
		smtp.PlainAuth("", cfg.MailFrom, cfg.MailPass, cfg.MailServer), cfg.MailFrom, []string{to}, msg)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Helper function help you bind data to the template
func parseTemplate(templatePath string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
