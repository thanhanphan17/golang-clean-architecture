package db

import (
	"go-clean-architecture/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	appConfig := config.AppConfig{
		BuildEnv:   DEV,
		DBHost:     "localhost",
		DBPort:     "5432",
		DBName:     "db_quizzer",
		DBUserName: "postgres",
		DBPassword: "0000",
	}

	db, err := NewDB(appConfig)
	assert.NoError(t, err)
	assert.NotNil(t, db)
}
