package database

import (
	"fmt"
	"os"
)

type Config struct {
	Protocol string
	Database string
	Username string
	Password string
	Host     string
	Port     string
}

func (c *Config) FillConfig() {
	c.Protocol = "postgresql"
	c.Database = os.Getenv("POSTGRES_DB")
	c.Username = os.Getenv("POSTGRES_USER")
	c.Password = os.Getenv("POSTGRES_PASSWORD")
	c.Host = os.Getenv("HOST")
	c.Port = os.Getenv("PORT")
}

func (c *Config) DatabaseUrl() string {
	database_url := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s", c.Protocol, c.Username, c.Password, c.Host, c.Port, c.Database)
	return database_url
}
