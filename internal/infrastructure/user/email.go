package infrastructure

import (
	"bytes"
	"main/internal/config"
	"main/internal/domain/user"
	"main/pkg"
	"path/filepath"
	"strconv"
	"text/template"

	"gopkg.in/gomail.v2"
)

const (
	AccountActivation = "Account activation"
)

type GmailService struct {
	logger             pkg.Logger
	fromEmail          string
	appPass            string
	appDomain          string
	smtp               string
	port               int
	apiActivationRoute string
	templatesDir       string
}

type EmailTemplateData struct {
	ActivationLink string
}

func NewGmailService(env config.Env, logger pkg.Logger) user.EmailService {
	port, _ := strconv.Atoi(env.GmailPort)
	return &GmailService{
		apiActivationRoute: env.ApiActivationRoute,
		fromEmail:          env.GmailFrom,
		appPass:            env.GmailPass,
		appDomain:          env.AppDomain,
		smtp:               env.GmailSMTP,
		templatesDir:       env.TemplatesDir,
		port:               port,
		logger:             logger,
	}
}

func (s *GmailService) SendToken(token string, email string) error {
	activationLink := s.appDomain + s.apiActivationRoute + "?token=" + token

	htmlBody, err := s.renderTemplate("activation.html", EmailTemplateData{
		ActivationLink: activationLink,
	})
	if err != nil {
		return err
	}

	if err := s.sendEmail(AccountActivation, email, htmlBody); err != nil {
		return err
	}

	return nil
}

func (s *GmailService) sendEmail(subject string, to string, htmlBody string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "NoReply <"+s.fromEmail+">")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)

	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer(s.smtp, s.port, s.fromEmail, s.appPass)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func (s *GmailService) renderTemplate(templateName string, data any) (string, error) {
	templatePath := filepath.Join(s.templatesDir, templateName)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
