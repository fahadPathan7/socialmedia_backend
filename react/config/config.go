package config

// Config struct def
type Config struct {
	Database                 string
	ReactColl                string
}

// New returns a new config
func New() Config {
	return Config{
		Database:                 "react",
		ReactColl:                "react",
	}
}
