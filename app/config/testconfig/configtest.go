package testconfig

import (
	"database/sql"
	"fmt"
	"go/build"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/jtorz/phoenix-backend/app/config"
	"github.com/jtorz/phoenix-backend/app/config/serverconfig"
	"github.com/spf13/viper"

	// postgres driver
	_ "github.com/lib/pq"
)

var mainDB *sql.DB

func init() {
	var err error
	viper.SetEnvPrefix(config.EnvPrefix)
	viper.SetTypeByDefaultValue(true)
	viper.BindEnv("DB_MAIN_CONNECTION")

	if conn := viper.GetString("DB_MAIN_CONNECTION"); conn != "" {
		mainDB, err = sql.Open("postgres", conn)
		if err != nil {
			log.Printf("Could't connect to %s_DB_MAIN_CONNECTION database.", config.EnvPrefix)
		}
	} else {
		log.Printf("no %s_DB_MAIN_CONNECTION database provided.", config.EnvPrefix)
	}
}

func MainDB() *sql.DB {
	return mainDB
}

type Config struct {
	DB           *sql.DB
	Redis        *redis.Pool
	Port         int
	Protocol     string
	JWTKey       string
	AppPathEnv   string
	LoggingLevel config.LogginLvl
}
type envConfig struct {
	Port         int              `mapstructure:"PORT"`
	Protocol     string           `mapstructure:"PROTOCOL"`
	JWTKey       string           `mapstructure:"JWT_KEY"`
	AppPathEnv   string           `mapstructure:"PATH"`
	LoggingLevel config.LogginLvl `mapstructure:"LOGGING_LEVEL"`
	// Database
	DBConnection string `mapstructure:"DB_MAIN_CONNECTION"`
	// Redis
	RedisAddress  string `mapstructure:"REDIS_ADDRESS"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
}

func NewConfig() (*Config, error) {
	c, err := loadconfig()
	if err != nil {
		return nil, fmt.Errorf("error loading test config: %s", err)
	}

	db, err := serverconfig.OpenPostgres(c.DBConnection, 5, 5, 5, 5)
	if err != nil {
		return nil, fmt.Errorf("connecting database: %w", err)
	}
	redis, err := serverconfig.OpenRedis(c.RedisAddress, c.RedisPassword, 5, 5)
	if err != nil {
		return nil, fmt.Errorf("connecting redis: %w", err)
	}
	return &Config{
		DB:           db,
		Redis:        redis,
		Port:         c.Port,
		Protocol:     c.Protocol,
		JWTKey:       c.JWTKey,
		AppPathEnv:   c.AppPathEnv,
		LoggingLevel: c.LoggingLevel,
	}, nil
}

func loadconfig() (*envConfig, error) {
	conf := envConfig{}
	viper.SetEnvPrefix("TEST_" + config.EnvPrefix)
	viper.SetTypeByDefaultValue(true)
	serverconfig.RegisterEnvs(conf)

	viper.SetDefault("LOGGING_LEVEL", "debug")

	{ //PATH

		p := os.Getenv("GOPATH")
		if p == "" {
			p = build.Default.GOPATH
		}
		p = addSlash(p) + "src/" + config.SysPkgName + "/"
		viper.SetDefault("PATH", p)
	}

	err := viper.Unmarshal(&conf)
	if err != nil {
		return &conf, nil
	}
	return nil, err
}

func addSlash(s string) string {
	if s[len(s)-1] != '/' {
		return s + "/"
	}
	return s
}
