package config

import (
	"os"
	"strconv"
)

type Config struct {
	App      *App
	Database *Database
	Kafka    *Kafka
}

type App struct {
	Name    string
	Env     string
	Port    int
	LogFile string
	BinFile string
	JWTKey  string
}

type Database struct {
	Host        string
	Port        int
	User        string
	Password    string
	Name        string
	SSLMode     string
	MaxIdle     int
	MaxOpen     int
	MaxLifeTime int
	MaxIdleTime int
}

type Kafka struct {
	Host    string
	Port    int
	Topic   string
	GroupID string
}

func LoadConfig() *Config {
	return &Config{
		App:      LoadAppConfig(),
		Database: LoadDatabaseConfig(),
		Kafka:    LoadKafkaConfig(),
	}
}

func LoadAppConfig() *App {
	return &App{
		Name:    getEnv("APP_NAME", "go-kafka"),
		Env:     getEnv("APP_ENV", "development"),
		Port:    getEnvAsInt("APP_PORT", 50051),
		LogFile: getEnv("APP_LOG_FILE", "./logs/go-kafka.log"),
		BinFile: getEnv("APP_BIN_FILE", "./bin/go-kafka"),
		JWTKey:  getEnv("APP_JWT_KEY", "secret"),
	}
}

func LoadDatabaseConfig() *Database {
	return &Database{
		Host:        getEnv("DB_HOST", "localhost"),
		Port:        getEnvAsInt("DB_PORT", 5432),
		User:        getEnv("DB_USER", "postgres"),
		Password:    getEnv("DB_PASSWORD", "secret"),
		Name:        getEnv("DB_NAME", "go_kafka"),
		SSLMode:     getEnv("DB_SSL_MODE", "disable"),
		MaxIdle:     getEnvAsInt("DB_MAX_IDLE", 10),
		MaxOpen:     getEnvAsInt("DB_MAX_OPEN", 100),
		MaxLifeTime: getEnvAsInt("DB_MAX_LIFE_TIME", 300),
		MaxIdleTime: getEnvAsInt("DB_MAX_IDLE_TIME", 300),
	}
}

func LoadKafkaConfig() *Kafka {
	return &Kafka{
		Host:    getEnv("KAFKA_HOST", "localhost"),
		Port:    getEnvAsInt("KAFKA_PORT", 9092),
		Topic:   getEnv("KAFKA_TOPIC_RECONCILE_USER", "reconcile-user-topic"),
		GroupID: getEnv("KAFKA_GROUP_ID", "go-kafka-group-id"),
	}
}

func GetAppPort() string {
	return getEnv("APP_PORT", "50051")
}

func GetAppEnv() string {
	return getEnv("APP_ENV", "development")
}

func GetAppBinFile() string {
	return getEnv("APP_BIN_FILE", "./bin/go-kafka")
}

func GetAppJWTKey() string {
	return getEnv("APP_JWT_KEY", "secret")
}

// getEnv returns the value of the environment variable or a default value if not set
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

// getEnvAsInt returns the value of the environment variable as an integer or a default value if not set
func getEnvAsInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return defaultVal
}
