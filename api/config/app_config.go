package config

var (
	// App application details
	App = struct {
		Name    string
		Usage   string
		Version string
	}{
		Name:    "storeService",
		Usage:   "API for interacting with the store service",
		Version: "0.0.1",
	}

	// Prefix configuration prefix
	Prefix = "APP"

	// DefaultMigrationDirectory db migration path
	DefaultMigrationDirectory = "db/migrate"
)
