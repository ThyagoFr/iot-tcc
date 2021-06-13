package database

import (
	"fmt"

	"github.com/thyagofr/tcc/api/infra/env"
)

// Config struct hold database info to create a new connection
type Config struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

func (c *Config) URL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Database)
}

// DNS used for database connection.
func (c *Config) DNS() string {
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable port=%d",
		c.Host,
		c.Database,
		c.User,
		c.Password,
		c.Port,
	)
}

// ConfigFromEnv configure using environmnet variables
func ConfigFromEnv(conf *Config) {
	conf.Host = env.GetOrDefaultString("DB_HOST", "localhost")
	conf.Port = env.GetOrDefaultInt("DB_PORT", 27017)
	conf.Database = env.GetOrDefaultString("DB_NAME", "tcc")
	conf.User = env.GetOrDefaultString("DB_USER", "tcc")
	conf.Password = env.GetOrDefaultString("DB_PASSWORD", "tcc")
}
