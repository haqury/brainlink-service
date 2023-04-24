package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"runtime"
	"strconv"
)

const (
	defaultHttpPort = ":80"
	defaultGrpcPort = ":81"
	AppEnvLocal     = "local"
	AppEnvTest      = "test"
	AppEnvProd      = "production"
)

// Config represents an application configuration.
type Config struct {
	AppEnv             string
	HttpPort           string
	GrpcPort           string
	DSN                string
	DbSchema           string
	MigrationPath      string
	SwaggerPath        string
	ExternalServiceUrl string
	WorkerPoolSize     int
	JaegerHost         string
}

func Get() *Config {
	if err := godotenv.Load(); err != nil {
		if check := os.IsNotExist(err); !check {
			fmt.Printf("failed to load env vars: %s", err)
		}
	}

	env := getEnv("APP_ENV", AppEnvLocal)

	dsn := fmt.Sprintf(
		"host=%s port=%s database=%s user=%s password=%s",
		getEnv("DB_HOST", ""),
		getEnv("DB_PORT", ""),
		getEnv("DB_DATABASE", ""),
		getEnv("DB_USER", ""),
		getEnv("DB_PASSWORD", ""),
	)

	if env == AppEnvLocal {
		dsn = dsn + " sslmode=disable"
	}

	return &Config{
		AppEnv:             env,
		HttpPort:           getEnv("HTTP_PORT", defaultHttpPort),
		GrpcPort:           getEnv("GRPC_PORT", defaultGrpcPort),
		DSN:                dsn,
		DbSchema:           getEnv("DB_SCHEMA", ""),
		SwaggerPath:        getEnv("SWAGGER_PATH", "api/api.swagger.json"),
		ExternalServiceUrl: getEnv("EXTERNAL_SERVICE_URL", ""),
		WorkerPoolSize:     getEnvAsInt("WORKER_POOL_SIZE", runtime.NumCPU()),
		JaegerHost:         getEnv("JAEGER_HOST", ""),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	value := getEnv(key, "")
	if v, e := strconv.Atoi(value); e == nil {
		return v
	}

	return defaultValue
}
