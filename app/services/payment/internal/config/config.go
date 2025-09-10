package config

import (
	"fmt"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName      string `envconfig:"APP_NAME" default:"user-service"`
	AppPort      string `envconfig:"APP_PORT" default:":8080"`
	StripeSecret string `envconfig:"STRIPE_SECRET"`
	LogLevel     string `envconfig:"LOG_LEVEL" default:"info"`
	JWTSecret    string `envconfig:"JWT_SECRET" default:"secret"`
	Domain       string `envconfig:"DOMAIN" default:"http://localhost"`
	Databases    struct {
		PostgresDSN string `envconfig:"POSTGRES_DSN"`
	}

	GrpcPort string `envconfig:"GRPC_PORT" default:":4040"`

	MessageBroker struct {
		Enabled bool   `envconfig:"MESSAGE_BROKER_ENABLED" default:"false"`
		Type    string `envconfig:"MESSAGE_BROKER_TYPE" default:"kafka"`
		URL     string `envconfig:"MESSAGE_BROKER_URL"`
		Topic   string `envconfig:"MESSAGE_BROKER_TOPIC"`
	}
}

var (
	cfg  *Config
	err  error
	once sync.Once
)

func GetConfig(envfiles ...string) (*Config, error) {
	once.Do(func() {
		_ = godotenv.Load(envfiles...)
		var c Config
		if err = envconfig.Process("", &c); err != nil {
			err = fmt.Errorf("failed to process envconfig: %v", err)
			return
		}
		cfg = &c
	})

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
