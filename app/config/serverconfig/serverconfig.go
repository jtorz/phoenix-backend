package serverconfig

import (
	"database/sql"
	"go/build"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/jtorz/phoenix-backend/app/config"
	"github.com/spf13/viper"
)

type Config struct {
	Port           int              `mapstructure:"PORT"`
	PortRedirect   int              `mapstructure:"PORT_REDIRECT"`
	Domain         string           `mapstructure:"DOMAIN"`
	Protocol       string           `mapstructure:"PROTOCOL"`
	RequestTimeout int              `mapstructure:"REQUEST_TIMEOUT"`
	Cert           string           `mapstructure:"SERVER_CERT"`
	Key            string           `mapstructure:"SERVER_KEY"`
	AppMode        string           `mapstructure:"MODE"`
	JWTKey         string           `mapstructure:"JWT_KEY"`
	CryptKey       string           `mapstructure:"CRYPT_KEY"`
	AppPathEnv     string           `mapstructure:"PATH"`
	LoggingLevel   config.LogginLvl `mapstructure:"LOGGING_LEVEL"`
	// Database
	DBMainConnection      string `mapstructure:"DB_MAIN_CONNECTION"`
	DBMainMaxIdleConns    int    `mapstructure:"DB_MAIN_MAX_IDLE_CONNS"`
	DBMainMaxOpenConns    int    `mapstructure:"DB_MAIN_MAX_OPEN_CONNS"`
	DBMainConnMaxIdleTime int    `mapstructure:"DB_MAIN_CONN_MAX_IDLE_TIME"`
	DBMainConnMaxLifetime int    `mapstructure:"DB_MAIN_CONN_MAX_LIFETIME"`

	// Redis
	RedisMaxIdleConns int    `mapstructure:"REDIS_MAX_IDLE_CONNS"`
	RedisMaxOpenConns int    `mapstructure:"REDIS_MAX_OPEN_CONNS"`
	RedisAddress      string `mapstructure:"REDIS_ADDRESS"`
	RedisPassword     string `mapstructure:"REDIS_PASSWORD"`
}

func (c Config) AppModeDebug() bool {
	return c.AppMode == "" || c.AppMode == "debug" || c.AppMode == "qa"
}

func LoadConfig() (*Config, error) {
	conf := Config{}
	viper.SetEnvPrefix(config.EnvPrefix)
	viper.SetTypeByDefaultValue(true)
	RegisterEnvs(conf)

	viper.SetDefault("REQUEST_TIMEOUT", 10)
	viper.SetDefault("DB_MAIN_MAX_IDLE_CONNS", 10)
	viper.SetDefault("DB_MAIN_MAX_OPEN_CONNS", 5)
	viper.SetDefault("DB_MAIN_CONN_MAX_IDLE_TIME", 300)
	viper.SetDefault("DB_MAIN_CONN_MAX_LIFETIME", 600)
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
		return nil, err
	}
	return &conf, nil
}

func RegisterEnvs(iv interface{}) {
	v := reflect.ValueOf(iv)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("mapstructure")
		tagValues := strings.Split(tag, ",")
		viper.BindEnv(tagValues[0])
	}
}

func OpenPostgres(connection string, maxOpen, maxIdle, maxLifeTime, maxIdleTime int) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)
	db.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(maxIdleTime))

	var one int
	err = db.QueryRow("SELECT 1").Scan(&one)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func OpenRedis(address, pass string, maxOpen, maxIdle int) (*redis.Pool, error) {
	redis := redis.Pool{
		MaxIdle:     maxOpen,
		MaxActive:   maxIdle,
		IdleTimeout: 240 * time.Second,
		TestOnBorrow: func(c redis.Conn, _ time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address)
			if err != nil {
				return nil, err
			}
			if pass != "" {
				if _, err := c.Do("AUTH", pass); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}

	conn := redis.Get()
	defer conn.Close()
	_, err := conn.Do("PING")
	if err != nil {
		return nil, err
	}
	return &redis, nil
}

func addSlash(s string) string {
	if s[len(s)-1] != '/' {
		return s + "/"
	}
	return s
}
