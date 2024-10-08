package config

import (
    "os"
)

type DBConfig struct {
    User     string
    Password string
    Host     string
    Port     string
    DBName   string
}

type ServerConfig struct {
    Port string
}

func LoadConfig() (*DBConfig, *ServerConfig) {
    dbConfig := &DBConfig{
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        DBName:   os.Getenv("DB_NAME"),
    }

    serverConfig := &ServerConfig{
        Port: os.Getenv("SERVER_PORT"),
    }

    return dbConfig, serverConfig
}
