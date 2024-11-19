package smtp

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

// SMTPClient contiene la configuración necesaria para enviar correos vía SMTP
type SMTPClient struct {
	Host     string
	Port     string
	Email    string
	password string // Minuscula para no mostrar la contraseña en el código
}

// NewSMTPClient crea una nueva instancia de SMTPClient
func NewSMTPClient(host, port, email, password string) *SMTPClient {
	log.Println("Conexión a SMTP establecida")
	return &SMTPClient{
		Host:     host,
		Port:     port,
		Email:    email,
		password: password,
	}
}

// Send implementa el método Send de la interfaz EmailSender para SMTPClient
func (s *SMTPClient) Send(to, subject, body string) error {
	e := email.NewEmail()
	e.From = s.Email
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(body)

	// Conexión SMTP con TLS
	address := fmt.Sprintf("%s:%s", s.Host, s.Port)
	err := e.Send(address, smtp.PlainAuth("", s.Email, s.password, s.Host))
	if err != nil {
		return fmt.Errorf("error enviando correo: %w", err)
	}

	return nil
}
