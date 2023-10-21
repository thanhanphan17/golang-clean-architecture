package db

import (
	"fmt"
	"go-clean-architecture/config"
	"go-clean-architecture/migration"

	"log/slog"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	PROD = "prod"
	DEV  = "dev"
)

type Database struct {
	Executor *gorm.DB
}

// InitDatabase initializes the database connection and performs optional upgrades.
// It returns a pointer to the Database struct.
func InitDatabase(allowUpgrade bool, serviceConfig config.AppConfig) *Database {
	// Log starting connection
	slog.Info("Starting connect database...")

	// Create a new database instance
	db, err := NewDB(serviceConfig)
	if err != nil {
		panic(err)
	}

	// Perform database migration if allowed
	if allowUpgrade {
		err := migration.Migration(db.Executor)
		if err != nil {
			return nil
		}
	}

	// Log successful connection
	slog.Info("Database connected!")
	return db
}

// NewDB creates a new database connection based on the provided configuration.
// It returns a pointer to the Database struct and an error if any.
func NewDB(config config.AppConfig) (*Database, error) {
	// Set the SSL mode based on the build environment
	// var configSSLMode = Disable
	// if config.BuildEnv == PROD {
	// 	configSSLMode = Require
	// }

	// Create the connection configuration
	cfg := Connection{
		SSLMode:                     Disable,
		Host:                        config.DBHost,
		Port:                        config.DBPort,
		Database:                    config.DBName,
		User:                        config.DBUserName,
		Password:                    config.DBPassword,
		SSLCertAuthorityCertificate: config.SSLCertAuthorityCertificate,
		MaxOpenConnections:          config.MaxOpenConnections,
		MaxIdleConnections:          config.MaxIdleConnections,
		ConnectionMaxIdleTime:       time.Duration(config.ConnectionMaxIdleTime),
		ConnectionMaxLifeTime:       time.Duration(config.ConnectionMaxLifeTime),
		ConnectionTimeout:           time.Duration(config.ConnectionTimeout),
	}

	// Log the generated Postgres connection string
	slog.Info(cfg.ToPostgresConnectionString())

	// Open a database connection using the gorm library
	db, err := gorm.Open(postgres.Open(cfg.ToPostgresConnectionString()), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	// Get the underlying *sql.DB instance from the gorm.DB object
	settingDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	// Ping the database to check the connection
	if pingError := settingDb.Ping(); pingError != nil {
		panic(pingError)
	}

	// Log the successful connection
	slog.Info(fmt.Sprintf("Connected to database: %s", config.DBName))

	// Configure the connection pool settings
	settingDb.SetMaxOpenConns(cfg.MaxOpenConnections)
	settingDb.SetMaxIdleConns(cfg.MaxIdleConnections)
	settingDb.SetConnMaxIdleTime(cfg.ConnectionMaxIdleTime)
	settingDb.SetConnMaxLifetime(cfg.ConnectionMaxLifeTime)

	// Return the Database struct with the gorm.DB object
	return &Database{
		Executor: db,
	}, nil
}
