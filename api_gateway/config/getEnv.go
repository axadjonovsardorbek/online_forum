package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT string

	FORUM_SERVICE_PORT string

	LOG_PATH string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":8081"))

	config.LOG_PATH = cast.ToString(coalesce("LOG_PATH", "logs/info.log"))
	config.FORUM_SERVICE_PORT = cast.ToString(coalesce("FORUM_SERVICE_PORT", ":50051"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
