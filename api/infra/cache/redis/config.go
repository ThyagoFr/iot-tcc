package redis

import (
	"fmt"

	"github.com/thyagofr/tcc/api/infra/env"
)

// Config Redis configuration
type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

type ConfigOpt func(*Config) error

func ConfigFromEnv(config *Config) error {
	config.Host = env.GetOrDefaultString("REDIS_HOST", "localhost")
	config.Port = env.GetOrDefaultInt("REDIS_PORT", 6379)
	config.Database = env.GetOrDefaultInt("REDIS_DB", 0)
	return nil
}

// DNS used for database connection.
func (c *Config) DNS() string {
	return fmt.Sprintf("%s:%d",
		c.Host,
		c.Port,
	)
}
