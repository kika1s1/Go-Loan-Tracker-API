package email

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/kika1s1/Go-Loan-Tracker-API/internal/config"
)

type SimpleEmailSender struct {
	smtpHost     string
	smtpPort     string
	senderEmail  string
	senderPasswd string
}

func NewSimpleEmailSender(smtpHost, smtpPort, senderEmail, senderPasswd string) *SimpleEmailSender {
	return &SimpleEmailSender{
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
		senderEmail:  senderEmail,
		senderPasswd: senderPasswd,
	}
}

func (s *SimpleEmailSender) SendVerificationEmail(userEmail string, token string) error {
	config, err := config.Load()
	if err != nil {
		return err
	}
	url := config.APP_DOMAIN

	subject := "Verify Your Email"
	html, err := s.loadHTML("pkg/email/templates/verify_email.html")
	if err != nil {
		return err
	}
	html = strings.Replace(html, "{{token}}", token, -1)
	html = strings.Replace(html, "{{url}}", "http://"+url, -1)
	// fmt.Println(html)
	return s.sendMail(userEmail, subject, html)
}

func (s *SimpleEmailSender) SendPasswordResetEmail(userEmail string, token string) error {
	config, err := config.Load()
	if err != nil {
		return err
	}
	url := config.APP_DOMAIN

	subject := "Reset Your Password"
	html, err := s.loadHTML("pkg/email/templates/password_reset.html")
	html = strings.Replace(html, "{{url}}", "http://"+url, -1)
	if err != nil {
		return err
	}
	html = strings.Replace(html, "{{token}}", token, -1)
	return s.sendMail(userEmail, subject, html)
}

func (s *SimpleEmailSender) loadHTML(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *SimpleEmailSender) sendMail(to string, subject string, html string) error {
	auth := smtp.PlainAuth("", s.senderEmail, s.senderPasswd, s.smtpHost)
	from := s.senderEmail
	toList := []string{to}

	header := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: multipart/alternative; boundary=\"boundary\"\r\n\r\n", from, toList[0], subject)
	body := fmt.Sprintf("--boundary\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s\r\n--boundary--", html)

	msg := header + body

	// Debug output
	fmt.Printf("Sending email to: %s\n", to)
	fmt.Printf("Message:\n%s\n", msg)

	err := smtp.SendMail(s.smtpHost+":"+s.smtpPort, auth, from, toList, []byte(msg))
	if err != nil {
		fmt.Printf("Failed to send email: %v\n", err)
	}
	return err
}

func (s *SimpleEmailSender) SendPromotionToAdminEmail(userEmail string) error {
	config, err := config.Load()
	if err != nil {
		return err
	}
	url := config.APP_DOMAIN

	subject := "Congratulations! You've Been Promoted to Admin"
	html, err := s.loadHTML("pkg/email/templates/admin_promotion.html")
	if err != nil {
		return err
	}
	html = strings.Replace(html, "{{url}}", "http://"+url, -1)
	return s.sendMail(userEmail, subject, html)
}

func (s *SimpleEmailSender) SendDemotionFromAdminEmail(userEmail string) error {
	config, err := config.Load()
	if err != nil {
		return err
	}
	url := config.APP_DOMAIN

	subject := "Notice: You Have Been Demoted from Admin"
	html, err := s.loadHTML("pkg/email/templates/admin_demotion.html")
	if err != nil {
		return err
	}
	html = strings.Replace(html, "{{url}}", "http://"+url, -1)
	return s.sendMail(userEmail, subject, html)
}
