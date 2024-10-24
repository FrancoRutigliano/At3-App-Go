package service

import (
	"log"

	"github.com/wneessen/go-mail"
)

type EmailService struct {
	Mailer *mail.Client
	From   string
}

func (e *EmailService) New(smtpHost, smtpPort, from, password string) {
	// Crear un nuevo cliente SMTP
	client, err := mail.NewClient(smtpHost, mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(from), mail.WithPassword(password))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}
	e.Mailer = client

	e.From = from
}

func (e *EmailService) SendRegisterEmail(to, token string) error {
	message := mail.NewMsg()
	if err := message.From(e.From); err != nil {
		return err // Manejo de errores
	}
	if err := message.To(to); err != nil {
		return err // Manejo de errores
	}
	message.Subject("Confirma tu cuenta")

	// Cuerpo del correo en HTML
	body := generateEmailBody(token)
	message.SetBodyString(mail.TypeTextHTML, body)

	// Enviar el correo
	if err := e.Mailer.DialAndSend(message); err != nil {
		return err // Manejo de errores
	}
	return nil
}

func generateEmailBody(token string) string {
	return `<html>
        <head><title>Confirma tu cuenta</title></head>
        <body>
            <p>Gracias por registrarte. Por favor confirma tu cuenta haciendo clic en el siguiente enlace:</p>
            <a href="https://app.atomico3.io/confirm_account?token=` + token + `">Confirmar cuenta</a>
        </body>
    </html>`
}
