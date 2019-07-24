package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/jordan-wright/email"

	"github.com/lbrulet/APINIT-GO/configs"
	"github.com/lbrulet/APINIT-GO/models"
)

// SendMail is used to send an email
func SendMail(user models.User, formMail interface{}, pathTotemplate string) {
	if tmpl, err := template.ParseFiles(pathTotemplate); err != nil {
		fmt.Printf("[WARNING] %s!\n", err)
	} else {
		var tpl bytes.Buffer
		tmpl.Execute(&tpl, formMail)
		e := email.NewEmail()
		e.From = configs.Config.MailFrom
		e.To = []string{user.Email}
		e.Subject = configs.Config.MailSubject
		e.HTML = []byte(tpl.String())
		go func() {
			if len(configs.Config.MailAddress) != 0 && len(configs.Config.MailPassword) != 0 {
				e.Send(configs.Config.SMTPAddress+":"+configs.Config.SMTPPort, smtp.PlainAuth("", configs.Config.MailAddress,
					configs.Config.MailPassword, configs.Config.SMTPAddress))
			} else {
				fmt.Printf("[WARNING] You didn't load your credentials into the environnement!\n")
			}
		}()
	}
}
