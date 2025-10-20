package email

import (
	"fmt"

	"github.com/wneessen/go-mail"
)

type Mailer struct{}

func NewMailer() *Mailer {
	return &Mailer{}
}

func (s *Mailer) client() (*mail.Client, error) {
	client, err := mail.NewClient("host", mail.WithPort(8888), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername("m.cfg.Username"), mail.WithPassword("pwd"))
	if err != nil {
		return nil, err
	}
	client.SetSSL(true)
	return client, nil
}

// Send发送邮箱
func (s *Mailer) Send(msg ...*mail.Msg) error {
	client, err := s.client()
	if err != nil {
		return fmt.Errorf("mailer createClient error, %w", err)
	}
	err = client.DialAndSend(msg...)
	if err != nil {
		return fmt.Errorf("mailer send error, %w", err)
	}
	return nil
}
