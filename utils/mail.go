package utils

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
	"github.com/lbrulet/APINIT-GO/configs"
	"github.com/lbrulet/APINIT-GO/models"
)

type Template struct {
	Email        string
	Username     string
	ConfirmEmail string
}

func SendMail(user models.User, token string) {
	if pwd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		if tmpl, err := template.ParseFiles(pwd + "/templates/welcome.html"); err != nil {
		} else {
			data := Template{
				Email:        user.Email,
				Username:     user.Username,
				ConfirmEmail: configs.Config.MailConfirmationLink + "?token=" + token,
			}
			var tpl bytes.Buffer
			tmpl.Execute(&tpl, data)
			e := email.NewEmail()
			e.From = configs.Config.MailFrom
			e.To = []string{user.Email}
			e.Subject = configs.Config.MailSubject
			e.HTML = []byte(tpl.String())
			go func() {
				e.Send(configs.Config.SMTPAddress, smtp.PlainAuth("", "ADDRESSGOOGLE", "PASSWORD", "smtp.gmail.com"))
			}()
		}
	}
}
