package repository

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var testdb *sql.DB

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = 10 * time.Second
	if err != nil {
		log.Panicf("failed to connect to docker: %v", err)
	}

	pwd, _ := os.Getwd()

	runOptions := &dockertest.RunOptions{
		Repository: "timescale/timescaledb",
		Tag:        "latest-pg15",
		Env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=password",
			"POSTGRES_DB=testdb",
			"listen_addresses='*'",
		},
		Mounts: []string{
			pwd + "/../../schema.sql:/docker-entrypoint-initdb.d/schema.sql",
		},
	}

	resource, err := pool.RunWithOptions(runOptions,
		func(config *docker.HostConfig) {
			// config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)
	if err != nil {
		log.Panicf("failed to start resource: %v", err)
	}

	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("postgres", "password"),
		Host:   resource.GetHostPort("5432/tcp"),
		Path:   "testdb",
		RawQuery: url.Values{
			"sslmode": []string{"disable"},
		}.Encode(),
	}

	if err := pool.Retry(func() error {
		db, err := sql.Open("postgres", dsn.String())
		if err != nil {
			return fmt.Errorf("failed to make connection for postgres: %w", err)
		}

		if err := db.Ping(); err != nil {
			return fmt.Errorf("failed to ping for postgres: %w", err)
		}

		testdb = db

		return nil
	}); err != nil {
		log.Panicf("failed to connect to database: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("failed to purge resource: %s", err)
	}

	os.Exit(code)
}
