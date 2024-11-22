package config

// Config struct def
type Config struct {
	Database                 string
	PostColl                 string
}

// New returns a new config
func New() Config {
	return Config{
		Database:                 "post",
		PostColl:                 "post",
	}
}