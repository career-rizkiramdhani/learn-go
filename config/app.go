package config

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"

	"crud/utils"
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

	// JWT expiration in hours (integer)
	JWTExpiredHours int
}

var Cfg Config

// InitConfig reads environment variables (and .env file when present)
func InitConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, falling back to environment variables")
	}

	// Prefer full DATABASE_URL
	Cfg.DatabaseURL = utils.GetEnv("DATABASE_URL", "")
	if Cfg.DatabaseURL == "" {
		// Build DSN from components if DATABASE_URL not provided
		host := utils.GetEnv("POSTGRES_HOST", "localhost")
		port := utils.GetEnv("POSTGRES_PORT", "5432")
		user := utils.GetEnv("POSTGRES_USER", "postgres")
		pass := utils.GetEnv("POSTGRES_PASSWORD", "")
		dbname := utils.GetEnv("POSTGRES_DB", "appdb")
		sslmode := utils.GetEnv("POSTGRES_SSLMODE", "disable")

		Cfg.DatabaseURL = "host=" + host + " port=" + port + " user=" + user + " password=" + pass + " dbname=" + dbname + " sslmode=" + sslmode
	}

	Cfg.JWTSecret = utils.GetEnv("JWT_SECRET", "secret")
	Cfg.Port = utils.GetEnv("PORT", "8080")
	Cfg.UseTLS = utils.GetEnv("USE_TLS", "false")
	Cfg.CertFile = utils.GetEnv("CERT_FILE", "cert.pem")
	Cfg.KeyFile = utils.GetEnv("KEY_FILE", "key.pem")

	// JWT_EXPIRED: expiration in hours (default 72)
	if v := utils.GetEnv("JWT_EXPIRED", "72"); v != "" {
		if h, err := strconv.Atoi(v); err == nil {
			Cfg.JWTExpiredHours = h
		} else {
			Cfg.JWTExpiredHours = 72
		}
	} else {
		Cfg.JWTExpiredHours = 72
	}
}
