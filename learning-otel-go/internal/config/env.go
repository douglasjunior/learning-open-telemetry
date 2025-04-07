package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadEnv carrega as variáveis de ambiente do arquivo .env, se existir
func LoadEnv() {
	// Tenta carregar .env, mas não falha se o arquivo não existir
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	LogLevel string
	APIKey   string
}

type ServerConfig struct {
	Port string
	Env  string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func LoadConfig() *Config {
	// Tenta carregar as variáveis de ambiente do arquivo .env
	LoadEnv()

	port := getEnv("PORT", "8080")
	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	return &Config{
		Server: ServerConfig{
			Port: port,
			Env:  getEnv("ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     dbPort,
			User:     getEnv("DB_USER", "admin"),
			Password: getEnv("DB_PASS", "admin123"),
			Name:     getEnv("DB_NAME", "golearn_db"),
		},
		LogLevel: getEnv("LOG_LEVEL", "info"),
		APIKey:   getEnv("API_KEY", ""),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
