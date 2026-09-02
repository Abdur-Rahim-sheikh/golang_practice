package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecret   string
	// db info
	UserName string
	Password string
	Host     string
	Port     int
	DbName   string
}

var conf *Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("no .env file found, relying on environment variables")
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}
	service_name := os.Getenv("SERVICE_NAME")
	if service_name == "" {
		fmt.Println("service_name is required")
		os.Exit(1)
	}
	http_port := os.Getenv("HTTP_PORT")
	port, err := strconv.ParseInt(http_port, 10, 32)
	if http_port == "" || err != nil {
		fmt.Println("valid HTTP_PORT is required", err)
		os.Exit(1)
	}

	conf = &Config{
		Version:     version,
		ServiceName: service_name,
		HttpPort:    int(port),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}
}

func GetConfig() *Config {
	if conf == nil {
		LoadConfig()
	}
	return conf
}
