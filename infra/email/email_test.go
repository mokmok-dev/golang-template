package email

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"testing"
	"time"

	"github.com/mokmok-dev/golang-template/domain/configuration"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var testemailconfig configuration.Email

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = 10 * time.Second
	if err != nil {
		log.Panicf("failed to connect to docker: %v", err)
	}

	runOptions := &dockertest.RunOptions{
		Repository: "mailhog/mailhog",
		Tag:        "latest",
	}

	resource, err := pool.RunWithOptions(runOptions,
		func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)
	if err != nil {
		log.Panicf("failed to start resource: %v", err)
	}

	if err := pool.Retry(func() error {
		if _, err := smtp.Dial(resource.GetHostPort("1025/tcp")); err != nil {
			return fmt.Errorf("failed to dial SMTP server: %w", err)
		}

		resource.GetHostPort("1025/tcp")

		testemailconfig = configuration.Email{
			Host:   resource.GetBoundIP("1025/tcp"),
			Port:   resource.GetPort("1025/tcp"),
			Sender: "golang-template@mokmok.dev",
		}

		return nil
	}); err != nil {
		log.Panicf("failed to connect to SMTP server: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("failed to purge resource: %s", err)
	}

	os.Exit(code)
}
