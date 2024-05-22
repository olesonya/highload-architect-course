package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
)

type Config struct {
	envs *envs
}

type envs struct {
	LogLevel           string        `env:"LOG_LEVEL" env-default:"debug" env-description:"log level: trace, debug, info, warn, error, fatal, panic"`
	GRPCAddress        string        `env:"GRPC_ADDR" env-default:":50051" env-description:"IP:PORT to bind grpc"`
	HTTPAddress        string        `env:"HTTP_ADDR" env-default:":8080" env-description:"IP:PORT to bind http"`
	DBHost             string        `env:"DB_HOST" env-default:"127.0.0.1" env-description:"IP or hostname where DB resides"`
	DBPort             uint          `env:"DB_PORT" env-default:"5432" env-description:"DB's port"`
	DBName             string        `env:"DB_NAME" env-default:"postgres" env-description:"Database name"`
	DBUser             string        `env:"DB_USER" env-default:"postgres" env-description:"Username to connect to DB"`
	DBPassword         string        `env:"DB_PASSWORD" env-default:"postgres" env-description:"Password to connect to DB"`
	MigrationsPath     string        `env:"MIGRATIONS_PATH" env-default:"./scripts/database/" env-description:"Path where Migration SQL scripts reside"`
	MetricsBindAddress string        `env:"METRICS_BIND_ADDRESS" env-default:":2112" env-description:"IP:PORT to bind metrics HTTP socket"`
	ReconnectDBTimeout time.Duration `env:"RECONNECT_DB_TIMEOUT" env-default:"1m" env-description:"Reconnect to DB timeout"`
}

func New() (*Config, error) {
	e := new(envs)

	helpString, err := e.HelpString()
	if err != nil {
		logger.Fatalf("getting help string of env settings failed: %v", err)
	}

	logger.Info(helpString)

	err = cleanenv.ReadConfig(".env", e)
	if errors.Is(err, fs.ErrNotExist) {
		err = cleanenv.ReadEnv(e)
	}

	if err != nil {
		logger.Fatalf("read env config failed: %v", err)
	}

	return &Config{
		envs: e,
	}, nil
}

func (e *envs) HelpString() (string, error) {
	helpString, err := cleanenv.GetDescription(e, nil)
	if err != nil {
		return "", fmt.Errorf("get help string failed: %w", err)
	}

	return helpString, nil
}

func (c *Config) PrintDebug() {
	conf, err := json.MarshalIndent(c.envs, "", "  ")
	if err != nil {
		logger.Errorf("marshaling env config failed: %v", err)

		return
	}

	logger.Debugf("Config:\n%s", conf)
}

func (c *Config) GetLogLevel() logger.Level {
	level, err := logger.ParseLevel(c.envs.LogLevel)
	if err != nil {
		logger.Error(err)

		return logger.InfoLevel
	}

	return level
}

func (c *Config) GetGRPCAddress() string {
	return c.envs.GRPCAddress
}

func (c *Config) GetHTTPAddress() string {
	return c.envs.HTTPAddress
}

func (c *Config) GetDBHost() string {
	return c.envs.DBHost
}

func (c *Config) GetDBPort() uint {
	return c.envs.DBPort
}

func (c *Config) GetDBName() string {
	return c.envs.DBName
}

func (c *Config) GetDBUser() string {
	return c.envs.DBUser
}

func (c *Config) GetDBPassword() string {
	return c.envs.DBPassword
}

func (c *Config) GetMigrationsPath() string {
	return c.envs.MigrationsPath
}

func (c *Config) GetMetricsBindAddress() string {
	return c.envs.MetricsBindAddress
}

func (c *Config) GetReconnectDBTimeout() time.Duration {
	return c.envs.ReconnectDBTimeout
}
