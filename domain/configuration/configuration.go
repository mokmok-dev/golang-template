package configuration

type Config struct {
	Log      Log
	Server   Server
	Database Database
	Email    Email
	GCP      GCP
}

type Log struct {
	Level string `envconfig:"LOG_LEVEL" default:"info"`
}

type Server struct {
	Port string `envconfig:"SERVER_PORT" default:"8080"`
}

type Database struct {
	Host     string `envconfig:"DATABASE_HOST" default:"localhost"`
	Port     string `envconfig:"DATABASE_PORT" default:"5432"`
	User     string `envconfig:"DATABASE_USER" default:"postgres"`
	Password string `envconfig:"DATABASE_PASSWORD" default:"postgres"`
	Name     string `envconfig:"DATABASE_NAME" default:"golang-template"`
}

type Email struct {
	Host     string `envconfig:"EMAIL_HOST" default:"localhsot"`
	Port     string `envconfig:"EMAIL_PORT" default:"1025"`
	User     string `envconfig:"EMAIL_USER" default:"golang-template"`
	Password string `envconfig:"EMAIL_PASSWORD" default:"password"`
	Sender   string `envconfig:"EMAIL_SENDER" default:"sender@mokmok.dev"`
}

type GCP struct {
	ProjectID string `envconfig:"GCP_PROJECT_ID" default:"mokmok-dev"`
}
