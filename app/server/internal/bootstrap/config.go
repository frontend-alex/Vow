package bootstrap

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Env              string
	HTTPAddr         string
	DatabaseURL      string
	AllowedOrigins   []string
	RateLimitRPM     int
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	IdleTimeout      time.Duration
	ShutdownTimeout  time.Duration
	JWTAccessSecret  string
	JWTRefreshSecret string
	JWTAccessTTL     time.Duration
	JWTRefreshTTL    time.Duration
}

func LoadConfig() Config {
	loadDotEnv(".env")

	return Config{
		Env:              getEnv("APP_ENV", "development"),
		HTTPAddr:         getEnv("HTTP_ADDR", ":8080"),
		DatabaseURL:      os.Getenv("DATABASE_URL"),
		AllowedOrigins:   splitCSV(getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000,http://localhost:5173")),
		RateLimitRPM:     getEnvInt("RATE_LIMIT_RPM", 120),
		ReadTimeout:      getEnvDuration("HTTP_READ_TIMEOUT", 5*time.Second),
		WriteTimeout:     getEnvDuration("HTTP_WRITE_TIMEOUT", 10*time.Second),
		IdleTimeout:      getEnvDuration("HTTP_IDLE_TIMEOUT", 60*time.Second),
		ShutdownTimeout:  getEnvDuration("HTTP_SHUTDOWN_TIMEOUT", 10*time.Second),
		JWTAccessSecret:  getEnv("JWT_ACCESS_SECRET", "change-me-access-secret"),
		JWTRefreshSecret: getEnv("JWT_REFRESH_SECRET", "change-me-refresh-secret"),
		JWTAccessTTL:     getEnvDuration("JWT_ACCESS_TTL", 15*time.Minute),
		JWTRefreshTTL:    getEnvDuration("JWT_REFRESH_TTL", 30*24*time.Hour),
	}
}

func loadDotEnv(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}

		key = strings.TrimSpace(key)
		value = strings.Trim(strings.TrimSpace(value), `"'`)
		if key == "" || os.Getenv(key) != "" {
			continue
		}

		_ = os.Setenv(key, value)
	}
}

func getEnv(key string, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func getEnvInt(key string, fallback int) int {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func splitCSV(value string) []string {
	parts := strings.Split(value, ",")
	items := make([]string, 0, len(parts))
	for _, part := range parts {
		item := strings.TrimSpace(part)
		if item != "" {
			items = append(items, item)
		}
	}
	return items
}
