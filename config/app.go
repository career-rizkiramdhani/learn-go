package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Database connection. Prefer DATABASE_URL (Postgres DSN). If not set,
	// InitConfig will build from POSTGRES_* variables.
	DatabaseURL string
	JWTSecret   string
	Port        string

	// TLS settings (optional). When USE_TLS=true and cert/key are provided,
	// main will start the server with TLS.
	UseTLS   string
	CertFile string
	KeyFile  string
}

var Cfg Config

// InitConfig reads environment variables (and .env file when present)
func InitConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, falling back to environment variables")
	}

	// Prefer full DATABASE_URL
	Cfg.DatabaseURL = getEnv("DATABASE_URL", "")
	if Cfg.DatabaseURL == "" {
		// Build DSN from components if DATABASE_URL not provided
		host := getEnv("POSTGRES_HOST", "localhost")
		port := getEnv("POSTGRES_PORT", "5432")
		user := getEnv("POSTGRES_USER", "postgres")
		pass := getEnv("POSTGRES_PASSWORD", "")
		dbname := getEnv("POSTGRES_DB", "appdb")
		sslmode := getEnv("POSTGRES_SSLMODE", "disable")

		Cfg.DatabaseURL = "host=" + host + " port=" + port + " user=" + user + " password=" + pass + " dbname=" + dbname + " sslmode=" + sslmode
	}

	Cfg.JWTSecret = getEnv("JWT_SECRET", "secret")
	Cfg.Port = getEnv("PORT", "8080")
	Cfg.UseTLS = getEnv("USE_TLS", "false")
	Cfg.CertFile = getEnv("CERT_FILE", "cert.pem")
	Cfg.KeyFile = getEnv("KEY_FILE", "key.pem")
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
