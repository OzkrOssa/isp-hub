package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		DB    *DB
		Redis *Redis
	}
	DB struct {
		Connection string
		Host       string
		Port       string
		Name       string
		User       string
		Password   string
	}
	Redis struct {
		Host     string
		Port     string
		Password string
	}
)

func New() (*Container, error) {
	if !strings.Contains(os.Getenv("ENV"), "prod") {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		Name:       os.Getenv("DB_NAME"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
	}

	redis := &Redis{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	return &Container{
		DB:    db,
		Redis: redis,
	}, nil
}
