package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    ServerAddress string
    DBUri         string
    JWTSecret     string
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
        return nil, err
    }

    cfg := &Config{
        ServerAddress: os.Getenv("SERVER_ADDRESS"),
        DBUri:         os.Getenv("DB_URI"),
        JWTSecret:     os.Getenv("JWT_SECRET"),
    }

    return cfg, nil
}
