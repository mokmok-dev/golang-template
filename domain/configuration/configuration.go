package configuration

type Config struct {
	Log      Log
	Database Database
	GCP      GCP
}

type Log struct {
	Level string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type GCP struct {
	ProjectID string
}
