package start

import (
	"../config"
	"github.com/russross/meddler"
	"strings"
)

// Configure the database
func ConfigureDatabase() {
	// Default application database.
	switch strings.ToLower(config.Database) {
		case "mysql":
			meddler.Default = meddler.MySQL
		case "sqlite":
			meddler.Default = meddler.SQLite
		case "postgresql":
			meddler.Default = meddler.PostgreSQL
	}

	// Set meddler database debug to our app setting
	meddler.Debug = config.Debug
}