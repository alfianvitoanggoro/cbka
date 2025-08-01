package db

import (
	"database/sql"
	"fmt"
	"go-kafka/internal/config"
	"go-kafka/pkg/logger"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var sqlDB *sql.DB

func ConnectDB(configDB *config.Database) (*gorm.DB, error) {
	dsn := builderDSN(configDB)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		errMessage := err.Error()
		logger.WriteLogToFile("error", "db.ConnectDB.gorm", err, &errMessage)
		return nil, err
	}

	sqlConn, err := db.DB()

	if err != nil {
		errMessage := err.Error()
		logger.WriteLogToFile("error", "db.ConnectDB.sql", err, &errMessage)
		return nil, err
	}

	sqlDB = sqlConn

	// Connection Pooling
	sqlDB.SetMaxIdleConns(configDB.MaxIdle)
	sqlDB.SetMaxOpenConns(configDB.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(configDB.MaxLifeTime) * time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Duration(configDB.MaxIdleTime) * time.Minute)

	logger.Info("✅ Database connection Successfully")
	return db, nil
}

func CloseDB() {
	if sqlDB != nil {
		if err := sqlDB.Close(); err != nil {
			logger.Warn(fmt.Sprintf("⚠️ Error closing DB: %v", err))
		} else {
			logger.Info("🔌 DB connection closed.")
		}
	}
}

func builderDSN(configDB *config.Database) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", configDB.Host, configDB.User, configDB.Password, configDB.Name, configDB.Port, configDB.SSLMode)
}
