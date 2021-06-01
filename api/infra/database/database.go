package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
  "fmt"
	"github.com/thyago/tcc/api-service/infra/env"
)

// Config struct hold database info to create a new connection
type Config struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

// DNS create a URI connection string
func (conf *Config) DNS() string {
	if conf.User == "" || conf.Password == "" {
		return fmt.Sprintf("mongodb://%s:%d/%s", conf.Host, conf.Port, conf.Database)
	}
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
}

// ConfigFromEnv configure using environmnet variables
func ConfigFromEnv(conf *Config) {
	conf.Host = env.GetOrDefaultString("MONGODB_HOST", "localhost")
	conf.Port = env.GetOrDefaultInt("MONGODB_PORT", 27017)
	conf.Database = env.GetOrDefaultString("MONGODB_DB", "tcc")
	conf.User = env.GetOrDefaultString("MONGODB_USER", "tcc")
	conf.Password = env.GetOrDefaultString("MONGODB_PASSWORD","tcc")
}

// NewMongoDBClient creates a new mongodb client
func NewMongoDBClient(conf *Config) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.DNS()))
	if err != nil {
		return nil, err
	}
	return client.Database(conf.Database), nil
}
