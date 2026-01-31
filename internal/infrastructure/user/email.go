package infrastructure

import (
	"main/internal/config"
	"main/internal/domain/user"

	"gopkg.in/gomail.v2"
)

type GmailService struct {
	fromEmail string
	appPass   string
	domain    string
}

func NewGmailService(env config.Env) user.EmailService {
	return &GmailService{
		fromEmail: env.GmailFrom,
		appPass:   env.GmailPass,
		domain:    env.AppDomain,
	}
}

func (s *GmailService) SendToken(token string, email string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "NoReply <"+s.fromEmail+">")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Account activation")
	activationLink := s.domain + "/verify?token=" + token
	htmlBody := s.getEmailBody(activationLink)
	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer("smtp.gmail.com", 587, s.fromEmail, s.appPass)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func (s *GmailService) getEmailBody(activationLink string) string {
	return `
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Подтверждение email</title>
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333;">
  <p>Привет!</p>
  <p>Пожалуйста, подтвердите ваш email, нажав на кнопку ниже:</p>
  
  <a href="` + activationLink + `" 
     style="display: inline-block; 
            padding: 12px 24px; 
            background-color: #4CAF50; 
            color: white; 
            text-decoration: none; 
            border-radius: 4px; 
            font-weight: bold; 
            margin: 16px 0;">
    Verify email
  </a>

  <hr style="margin: 32px 0; border: 0; border-top: 1px solid #eee;">
  <p style="color: #777; font-size: 12px;">
    Это письмо отправлено автоматически. Пожалуйста, не отвечайте на него.
  </p>
</body>
</html>`
}
