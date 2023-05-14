package configuration

type Config struct {
	Log Log
	GCP GCP
}

type Log struct {
	Level string
}

type GCP struct {
	ProjectID string
}
