package db

import (
	"go-clean-architecture/config"
	"go-clean-architecture/migration"
	"log/slog"

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

func InitDatabase(allowUpgrade bool, serviceConfig config.AppConfig) *Database {
	slog.Info("Starting connect database...")

	db, err := NewDB(serviceConfig)
	if err != nil {
		panic(err)
	}

	if allowUpgrade {
		migration.Migration(db.Executor)
	}

	slog.Info("Database connected!")
	return db
}

// NewDB --set=sslrootcert=path/to/your-ssl.crt
func NewDB(config config.AppConfig) (*Database, error) {
	var configSSLMode = Disable

	if config.BuildEnv == PROD {
		configSSLMode = Require
	}

	cfg := Connection{
		SSLMode:                     configSSLMode,
		Host:                        config.DBHost,
		Port:                        config.DBPort,
		Database:                    config.DBName,
		User:                        config.DBUserName,
		Password:                    config.DBPassword,
		SSLCertAuthorityCertificate: config.SSLCertAuthorityCertificate,
	}

	slog.Info(cfg.ToConnectionString())

	db, err := gorm.Open(postgres.Open(cfg.ToConnectionString()), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	settingDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	if pingError := settingDb.Ping(); pingError != nil {
		panic(pingError)
	}

	settingDb.SetMaxOpenConns(cfg.MaxOpenConnections)
	settingDb.SetMaxIdleConns(cfg.MaxIdleConnections)
	settingDb.SetConnMaxIdleTime(cfg.ConnectionMaxIdleTime)
	settingDb.SetConnMaxLifetime(cfg.ConnectionMaxLifeTime)

	return &Database{
		Executor: db,
	}, nil
}
