package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	HTTPAddr          string
	DatabaseURL       string
	JWTSecret         string
	LogLevel          string
	RateLimitRequests int
	CORSOrigins       []string
}

const (
	httpAddrEnv          = "HTTP_ADDR"
	databaseURLEnv       = "DATABASE_URL"
	jwtSecretEnv         = "JWT_SECRET"
	logLevelEnv          = "LOG_LEVEL"
	rateLimitRequestsEnv = "RATE_LIMIT_REQUESTS"
	corsOriginsEnv       = "CORS_ORIGINS"

	defaultHTTPAddr          = ":8080"
	defaultDatabaseURL       = "postgres://postgres:postgres@localhost:5432/vow?sslmode=disable"
	defaultJWTSecret         = "change-me"
	defaultLogLevel          = "info"
	defaultRateLimitRequests = 120
	defaultCORSOrigins       = "*"
)

func Load() Config {
	loadDotEnv(".env")

	return Config{
		HTTPAddr:          env(httpAddrEnv, defaultHTTPAddr),
		DatabaseURL:       env(databaseURLEnv, defaultDatabaseURL),
		JWTSecret:         env(jwtSecretEnv, defaultJWTSecret),
		LogLevel:          env(logLevelEnv, defaultLogLevel),
		RateLimitRequests: envInt(rateLimitRequestsEnv, defaultRateLimitRequests),
		CORSOrigins:       envList(corsOriginsEnv, defaultCORSOrigins),
	}
}

func env(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func envInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil || parsed <= 0 {
		return fallback
	}
	return parsed
}

func envList(key, fallback string) []string {
	value := env(key, fallback)
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

func loadDotEnv(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
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
