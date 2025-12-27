package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port           string
	DBPath         string
	MaxImages      int
	MaxImageSizeMB int
	UploadDir      string
}

func Load() *Config {
	return &Config{
		Port:           getEnv("PORT", "8080"),
		DBPath:         getEnv("DB_PATH", "./data/yearwrap.db"),
		MaxImages:      getEnvInt("MAX_IMAGES", 12),
		MaxImageSizeMB: getEnvInt("MAX_IMAGE_SIZE_MB", 8),
		UploadDir:      getEnv("UPLOAD_DIR", "./media"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}