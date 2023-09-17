package db

type SSLMode string

const (
	VerifyFull SSLMode = "verify-full"
	VerifyCA   SSLMode = "verify-ca"
	Require    SSLMode = "require"
	Prefer     SSLMode = "prefer"
	Allow      SSLMode = "allow"
	Disable    SSLMode = "disable"
)
