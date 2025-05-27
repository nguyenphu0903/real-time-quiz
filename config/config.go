package config

import (
	"os"
	"strconv"
)

type Config struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	Port          string
}

func Load() *Config {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	return &Config{
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       db,
		Port:          os.Getenv("PORT"),
	}
}
