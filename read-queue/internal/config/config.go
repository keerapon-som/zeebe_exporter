package config

import (
	"encoding/hex"
	"fmt"
	"log"
	"log/slog"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var config *Config

type Config struct {
	ServerMode string
	IsDebug    bool
	LogLevel   slog.Level
	LogFormat  string

	Aws        AWSConfig
	PostgresDB PostgresDBConfig
}

type PostgresDBConfig struct {
	MaxOpenConn               int
	MaxIdleConn               int
	ConnMaxLifeTimeTTL        *time.Duration
	Postgres_connectionstring string
}

type AWSConfig struct {
	Region      string
	Profile     string
	SQSEndpoint string
}

func GetConfig() *Config {
	if config != nil {
		return config
	}
	godotenv.Load()

	doInit()

	return config
}

func GetConfigWithFilename(envFileName string) *Config {

	if godotenv.Load(envFileName) == nil {
		goto INIT_CONFIG
	}
	if godotenv.Load(fmt.Sprintf("../%s", envFileName)) == nil {
		goto INIT_CONFIG
	}
	if godotenv.Load(fmt.Sprintf("../../%s", envFileName)) == nil {
		goto INIT_CONFIG
	}
	if godotenv.Load(fmt.Sprintf("../../../%s", envFileName)) == nil {
		goto INIT_CONFIG
	}

	if godotenv.Load(fmt.Sprintf("../../../../%s", envFileName)) == nil {
		goto INIT_CONFIG
	}
	if godotenv.Load(fmt.Sprintf("../../../../../%s", envFileName)) == nil {
		goto INIT_CONFIG
	}

	log.Fatalln("failed to load .env file")
INIT_CONFIG:

	doInit()
	return config
}

func doInit() {
	config = &Config{
		ServerMode: getEnvString("SERVER_MODE", ""), // local,dev, staging, prod

		IsDebug:   getEnvBool("DEBUG", false),
		LogLevel:  getEnvLogLevel("LOG_LEVEL", slog.LevelInfo),
		LogFormat: getEnvString("LOG_FORMAT", "JSON"),

		Aws: AWSConfig{
			Region:      getEnvString("AWS_REGION", ""),
			Profile:     getEnvString("AWS_PROFILE", ""),
			SQSEndpoint: getEnvString("AWS_SQS_ENDPOINT", ""),
		},

		PostgresDB: PostgresDBConfig{
			MaxOpenConn:               getEnvInt("POSTGRES_MAX_OPEN_CONN", 0),
			MaxIdleConn:               getEnvInt("POSTGRES_MAX_IDLE_CONN", 0),
			ConnMaxLifeTimeTTL:        getEnvDurationFromSecondsNullable("POSTGRES_CONN_MAX_LIFETIME_TTL", 0),
			Postgres_connectionstring: getEnvString("POSTGRES_CONNECTIONSTRING", ""),
		},
	}
}

func Init() {
	GetConfig()
}

//lint:ignore U1000 Ignore unused code, it may require in the future
func getEnvString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

//lint:ignore U1000 Ignore unused code, it may require in the future
func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return intValue
}

//lint:ignore U1000 Ignore unused code, it may require in the future
func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}

	return boolValue
}

//lint:ignore U1000 Ignore unused code, it may require in the future
func getEnvStringArray(key string, defaultValue []string) []string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	result := strings.Split(value, ",")
	for i := range result {
		result[i] = strings.TrimSpace(result[i])
	}

	return result
}

//lint:ignore U1000 Ignore unused code, it may require in the future
func getEnvDurationFromSeconds(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		return defaultValue
	}

	return time.Duration(intValue) * time.Second
}

//lint:ignore U1000 Ignore unused code, it may require in the future
func getEnvDurationFromSecondsNullable(key string, defaultValue time.Duration) *time.Duration {
	value := os.Getenv(key)
	if value == "" {
		if defaultValue == 0 {
			return nil
		} else {
			return &defaultValue
		}
	}

	intValue, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		return &defaultValue
	}

	result := time.Duration(intValue) * time.Second
	return &result
}

//lint:ignore U1000 Ignore unused code, it may require in the future
func getEnvBytes(key string, defaultValue []byte) []byte {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	decodedByteArray, err := hex.DecodeString(value)
	if err != nil {
		return defaultValue
	}

	return decodedByteArray
}

//lint:ignore U1000 Ignore unused code, it may require in the future
func getEnvLogLevel(key string, defaultValue slog.Level) slog.Level {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	switch value {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return defaultValue
	}
}

//lint:ignore U1000 Ignore unused code, it may require in the future
func joinPath(baseurl string, paths ...string) string {
	combined, err := url.JoinPath(baseurl, paths...)
	if err != nil {
		return ""
	}
	return combined
}
