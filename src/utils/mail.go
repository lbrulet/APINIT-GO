package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/jordan-wright/email"

	"github.com/lbrulet/APINIT-GO/src/configs"
	"github.com/lbrulet/APINIT-GO/src/models"
)

// SendMail is used to send an email
func SendMail(user *models.User, formMail interface{}, pathTotemplate string) {
	if tmpl, err := template.ParseFiles(pathTotemplate); err != nil {
		fmt.Printf("[WARNING] %s!\n", err)
	} else {
		var tpl bytes.Buffer
		err := tmpl.Execute(&tpl, formMail)
		if err != nil {
			fmt.Printf("[WARNING] %s\n", err.Error())
			return
		}
		e := email.NewEmail()
		e.From = configs.Config.MailFrom
		e.To = []string{user.Email}
		switch formMail.(type) {
		case models.Template:
			e.Subject = configs.Config.MailSubjectConfirmAccount
		case models.TemplateRecovery:
			e.Subject = configs.Config.MailSubjectRecoveryAccount
		}
		e.HTML = tpl.Bytes()
		go func() {
			if len(configs.Config.MailAddress) != 0 && len(configs.Config.MailPassword) != 0 {
				err := e.Send(configs.Config.SMTPAddress+":"+configs.Config.SMTPPort, smtp.PlainAuth("", configs.Config.MailAddress,
					configs.Config.MailPassword, configs.Config.SMTPAddress))
				if err != nil {
					fmt.Printf("[WARNING] %s\n", err.Error())
				}
			} else {
				fmt.Printf("[WARNING] You didn't load your credentials into the environnement!\n\texport MAIL_ADDRESS=your_email\n\texport MAIL_PASSWORD=your_password\n")
			}
		}()
	}
}
