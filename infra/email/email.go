package email

import (
	"bytes"
	"embed"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/google/wire"
	"github.com/mokmok-dev/golang-template/domain/configuration"
	domain "github.com/mokmok-dev/golang-template/domain/email"
	"github.com/mokmok-dev/golang-template/domain/logger"
	"github.com/mokmok-dev/golang-template/domain/tracer"
)

var _ domain.Email = (*Email)(nil)

var NewEmailSet = wire.NewSet(
	wire.Bind(new(domain.Email), new(*Email)),
	NewEmail,
)

//go:embed templates/*
var FS embed.FS

type Email struct {
	logger logger.Logger
	tracer tracer.Tracer
	config configuration.Email
	auth   smtp.Auth
	addr   string
}

func NewEmail(
	logger logger.Logger,
	tracer tracer.Tracer,
	config configuration.Email,
) (*Email, error) {
	auth := smtp.PlainAuth("", config.User, config.Password, config.Host)
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)

	return &Email{
		logger: logger,
		tracer: tracer,
		config: config,
		auth:   auth,
		addr:   addr,
	}, nil
}

type sendInput struct {
	To      []string
	Subject string
	Body    string
}

func (e *Email) send(input sendInput) error {
	msg := strings.Join([]string{
		fmt.Sprintf("From: %s", e.config.Sender),
		fmt.Sprintf("To: %s", strings.Join(input.To, ",")),
		fmt.Sprintf("Subject: %s", input.Subject),
		"MIME-Version: 1.0",
		"Content-Type: text/html; charset=\"UTF-8\"",
		"",
		input.Body,
	}, "\r\n")
	bytes := bytes.NewBufferString(msg).Bytes()

	if err := smtp.SendMail(e.addr, e.auth, e.config.Sender, input.To, bytes); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
