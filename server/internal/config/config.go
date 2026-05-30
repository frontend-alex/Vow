// Package config exposes environment-backed application configuration.
package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config contains all application configuration values.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Google   OAuthConfig
	Apple    OAuthConfig
	SMTP     SMTPConfig
}

// ServerConfig contains HTTP server settings.
type ServerConfig struct {
	HTTPAddr string
	AppEnv   string
}

// DatabaseConfig contains database connection settings.
type DatabaseConfig struct {
	URL string
}

// JWTConfig contains JSON Web Token settings.
type JWTConfig struct {
	SigningKey      string
	Issuer          string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

// OAuthConfig contains OAuth/OIDC provider settings.
type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

// SMTPConfig contains outbound email settings.
type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	accessTokenTTL, err := duration("JWT_ACCESS_TOKEN_TTL", 15*time.Minute)
	if err != nil {
		return nil, err
	}

	refreshTokenTTL, err := duration("JWT_REFRESH_TOKEN_TTL", 168*time.Hour)
	if err != nil {
		return nil, err
	}

	smtpPort, err := integer("SMTP_PORT", 587)
	if err != nil {
		return nil, err
	}

	return &Config{
		Server: ServerConfig{
			HTTPAddr: stringValue("HTTP_ADDR", ":8080"),
			AppEnv:   stringValue("APP_ENV", "development"),
		},
		Database: DatabaseConfig{
			URL: stringValue("DATABASE_URL", "postgres://vow:vow@localhost:5432/vow?sslmode=disable"),
		},
		JWT: JWTConfig{
			SigningKey:      stringValue("JWT_SIGNING_KEY", "change-me-use-a-long-random-secret"),
			Issuer:          stringValue("JWT_ISSUER", "vow-api"),
			AccessTokenTTL:  accessTokenTTL,
			RefreshTokenTTL: refreshTokenTTL,
		},
		Google: OAuthConfig{
			ClientID:     stringValue("GOOGLE_CLIENT_ID", ""),
			ClientSecret: stringValue("GOOGLE_CLIENT_SECRET", ""),
			RedirectURL:  stringValue("GOOGLE_REDIRECT_URL", "http://localhost:8080/auth/google/callback"),
		},
		Apple: OAuthConfig{
			ClientID:     stringValue("APPLE_CLIENT_ID", ""),
			ClientSecret: stringValue("APPLE_CLIENT_SECRET", ""),
			RedirectURL:  stringValue("APPLE_REDIRECT_URL", "http://localhost:8080/auth/apple/callback"),
		},
		SMTP: SMTPConfig{
			Host:     stringValue("SMTP_HOST", ""),
			Port:     smtpPort,
			Username: stringValue("SMTP_USERNAME", ""),
			Password: stringValue("SMTP_PASSWORD", ""),
			From:     stringValue("SMTP_FROM", "no-reply@example.com"),
		},
	}, nil
}

func stringValue(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func duration(key string, fallback time.Duration) (time.Duration, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("parse %s: %w", key, err)
	}
	return parsed, nil
}

func integer(key string, fallback int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("parse %s: %w", key, err)
	}
	return parsed, nil
}
