package config

// Config struct def
type Config struct {
	Database                 string
	UserColl                 string
	ForgottenPasswordRecords string
	UserDeletionRequestColl  string
	UserStoreColl            string
}

// New returns a new config
func New() Config {
	return Config{
		Database:                 "user",
		UserColl:                 "user",
		ForgottenPasswordRecords: "Forgotten",
		UserDeletionRequestColl:  "UserDeletionRequestColl",
		UserStoreColl:            "UserStoreColl",
	}
}
