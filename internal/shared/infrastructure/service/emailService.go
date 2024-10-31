package service

import (
	"fmt"
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
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="es">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Confirmación de Cuenta - Atomico3</title>
    <style>
      body {
        font-family: 'Arial', sans-serif;
        background: #1c1c1e !important;
        margin: 0;
        padding: 0;
        color: #fff !important;
      }
      p { color: #fff !important; }
      .container {
        max-width: 600px;
        margin: 40px auto;
        background-color: #2c2c2e;
        padding: 20px;
        border-radius: 8px;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.6);
      }
      h1 {
        color: #2f52ba;
        text-align: center;
      }
      .cta {
        margin-top: 30px;
        text-align: center;
      }
      .cta a {
        background-color: #2f52ba;
        color: white;
        text-decoration: none;
        padding: 12px 24px;
        border-radius: 8px;
        font-size: 16px;
      }
      .cta a:hover {
        background-color: #1e3a8a;
      }
      .footer {
        margin-top: 40px;
        text-align: center;
        color: #b0b0b0;
        font-size: 12px;
      }
      .footer a {
        color: #2f52ba;
        text-decoration: none;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <h1>Confirmación de Cuenta</h1>
      <p>Gracias por registrarte. Por favor confirma tu cuenta haciendo clic en el siguiente enlace:</p>
      <div class="cta">
        <a href="https://app.atomico3.io/confirm_account?token=%s">
          Confirmar cuenta
        </a>
      </div>
      <div class="footer">
        <p>Saludos cordiales,<br />El equipo de <strong>Atomico3</strong></p>
        <p>© 2024 Atomico3. Todos los derechos reservados.<br />
          <a href="https://atomico3.com/terms">Términos y Condiciones</a>
        </p>
      </div>
    </div>
  </body>
</html>`, token)
}
