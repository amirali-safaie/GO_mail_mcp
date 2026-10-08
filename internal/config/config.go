package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	IMAPHost string
	IMAPUser string
	IMAPPass string
	SMTPHost string
	SMTPPort int
}

func Load() Config {
	_ = godotenv.Load()

	port, _ := strconv.Atoi(getEnv("SMTP_PORT", "587"))

	return Config{
		IMAPHost: getEnv("MAIL_HOST", ""),
		IMAPUser: getEnv("MAIL_USER", ""),
		IMAPPass: getEnv("MAIL_PASS", ""),
		SMTPHost: getEnv("SMTP_HOST", ""),
		SMTPPort: port,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
