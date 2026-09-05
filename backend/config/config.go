package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_HOST     string
	POSTGRES_PORT     int
	POSTGRES_DB       string
}
type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecret   string
}

var conf *Config
var dbconfig *DBConfig

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
	postgres_port, err := strconv.ParseInt(os.Getenv("POSTGRES_PORT"), 10, 32)
	if err != nil {
		fmt.Println("postgres port needs to be valid int")
		os.Exit(1)
	}
	dbconfig = &DBConfig{
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:     int(postgres_port),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
	}
}

func GetConfig() *Config {
	if conf == nil {
		LoadConfig()
	}
	return conf
}

func GetDBConfig() *DBConfig {
	if dbconfig == nil {
		LoadConfig()
	}
	return dbconfig
}
