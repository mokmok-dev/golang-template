package configuration

type Config struct {
	Log      Log
	Server   Server
	Database Database
	Email    Email
	GCP      GCP
}

type Log struct {
	Level string
}

type Server struct {
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Email struct {
	Host     string
	Port     string
	User     string
	Password string
	Sender   string
}

type GCP struct {
	ProjectID string
}
