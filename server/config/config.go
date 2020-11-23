package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds configuration variables
type Config struct {
	Server ServerConfig
	MySQL  struct {
		Host     string
		Port     string
		Schema   string
		User     string
		Password string
	}
}

// ServerConfig holds server configuration variables
type ServerConfig struct {
	Host string
	Port string
	TLS  struct {
		Enabled  bool
		CertFile string // required if tls enabled
		KeyFile  string // required if tls enabled
	}
}

// MySQLConfig holds MySQL configuration variables
type MySQLConfig struct {
	Host     string
	Port     string
	Schema   string
	User     string
	Password string
}

// LoadConfig loads configuration variables from environment
func LoadConfig() *Config {
	godotenv.Load()

	config := &Config{}

	config.loadServerConfig()
	config.loadMySQLConfig()

	return config
}

func (c *Config) loadServerConfig() {
	c.Server.Host = os.Getenv("HOST")
	c.Server.Port = os.Getenv("PORT")
	c.Server.TLS.Enabled = os.Getenv("TLS_ENABLED") == "true"
	c.Server.TLS.CertFile = os.Getenv("TLS_CERT_FILE")
	c.Server.TLS.KeyFile = os.Getenv("TLS_KEY_FILE")
}

func (c *Config) loadMySQLConfig() {
	c.MySQL.Host = os.Getenv("MYSQL_HOST")
	c.MySQL.Port = os.Getenv("MYSQL_PORT")
	c.MySQL.Schema = os.Getenv("MYSQL_SCHEMA")
	c.MySQL.User = os.Getenv("MYSQL_USER")
	c.MySQL.Password = os.Getenv("MYSQL_PASSWORD")
}
