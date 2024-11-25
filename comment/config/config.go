package config

// Config struct def
type Config struct {
	Database                 string
	CommentColl              string
}

// New returns a new config
func New() Config {
	return Config{
		Database:                 "comment",
		CommentColl:              "comment",
	}
}