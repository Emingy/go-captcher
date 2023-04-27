package gocaptcher

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type Config struct {
	emojis []string
}

func (c *Config) initConfig() {
	// loads values from .env into the system
	if err := godotenv.Load("config.env"); err != nil {
		panic("No config.env file found")
	}

	emojis := getEnv("EMOJIS")
	if emojis != nil {
		c.emojis = strings.Split(*emojis, ",")
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string) *string {
	if value, exists := os.LookupEnv(key); exists {
		return &value
	}

	return nil
}
