package config

import (
	"github.com/caarlos0/env/v9"
	_ "github.com/joho/godotenv/autoload"
)

type AppConfig struct {
	Namespace      string `env:"NAMESPACE" envDefault:"example.com"`
	Debug          bool   `env:"DEBUG" envDefault:"false"`
	PhpPromosUrl   string `env:"PHP_PROMOS_URL"`
	GrpcPort       string `env:"GRPC_PORT" envDefault:"5050"`
	HttpPort       string `env:"HTTP_PORT" envDefault:"8888"`
	HttpCors       bool   `env:"HTTP_CORS" envDefault:"false"`
	WithMetrics    bool   `env:"WITH_METRICS" envDefault:"false"`
	WithTracing    bool   `env:"WITH_TRACING" envDefault:"false"`
	JaegerAddress  string `env:"JAEGER_ADDRESS"`
	RedisUrl       string `env:"REDIS_URL"`
	RedisPsw       string `env:"REDIS_PSW"`
	RedisDb        int    `env:"REDIS_DB"`
	OrdersFilePath string `env:"ORDERS_FILE_PATH"`
	PgDsn          string `env:"PG_DSN"`
	TestMode       bool   `env:"TEST_MODE" envDefault:"false"`
	BrokerUrl      string `env:"BROKER_URL"`
}

var Conf AppConfig

// UseDatabase checks if the application is configured to use a database.
func (c *AppConfig) UseDatabase() bool {
	return len(c.PgDsn) > 0
}

func init() {

	// Parse environment variables into the Conf struct
	if err := env.Parse(&Conf); err != nil {
		panic(err)
	}

}
