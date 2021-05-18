package server

import (
	"database/sql"
	"fmt"
	"go/build"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/kardianos/service"
	"github.com/spf13/viper"

	"github.com/jtorz/phoenix-backend/app/config"
	"github.com/jtorz/phoenix-backend/utils/syslog"

	// postgres driver
	_ "github.com/lib/pq"
)

type Server struct {
	Config     Config
	MainDB     *sql.DB
	HTTPServer *http.Server
	Services   map[string]Service
}

type Config struct {
	Port           int    `mapstructure:"PORT"`
	PortRedirect   int    `mapstructure:"PORT_REDIRECT"`
	Domain         string `mapstructure:"DOMAIN"`
	Protocol       string `mapstructure:"PROTOCOL"`
	RequestTimeout int    `mapstructure:"REQUEST_TIMEOUT"`
	Cert           string `mapstructure:"SERVER_CERT"`
	Key            string `mapstructure:"SERVER_KEY"`
	AppMode        string `mapstructure:"MODE"`
	JWTKey         string `mapstructure:"JWT_KEY"`
	CryptKey       string `mapstructure:"CRYPT_KEY"`
	AppPathEnv     string `mapstructure:"PATH"`
	// Database

	DBMainConnection      string `mapstructure:"DB_MAIN_CONNECTION"`
	DBMainMaxIdleConns    int    `mapstructure:"DB_MAIN_MAX_IDLE_CONNS"`
	DBMainMaxOpenConns    int    `mapstructure:"DB_MAIN_MAX_OPEN_CONNS"`
	DBMainConnMaxIdleTime int    `mapstructure:"DB_MAIN_CONN_MAX_IDLE_TIME"`
	DBMainConnMaxLifetime int    `mapstructure:"DB_MAIN_CONN_MAX_LIFETIME"`
}

func (c Config) AppModeDebug() bool {
	return c.AppMode == "" || c.AppMode == "debug" || c.AppMode == "qa"
}

func NewServer() Server {
	return Server{}
}

// Start should not block. Do the actual work async.
func (server Server) Start(s service.Service) error {
	log, err := s.Logger(nil)
	if err != nil {
		return err
	}
	err = server.load()
	if err != nil {
		log.Errorf("can't load server config: %s", err)
		return err
	}
	go server.runRedirect(log)
	log.Infof("listenAndServeTLS on port %d", server.Config.Port)
	if err := server.HTTPServer.ListenAndServeTLS(server.Config.Cert, server.Config.Key); err != nil && err != http.ErrServerClosed {
		log.Errorf("error ListenAndServeTLS: %s", err)
		os.Exit(1)
	}
	return nil
}

func (server Server) runRedirect(log syslog.Logger) {
	if server.Config.PortRedirect <= 0 {
		return
	}
	var origin, dest string
	{
		origin = fmt.Sprintf(":%d", server.Config.PortRedirect)
		dest = fmt.Sprintf("%s://%s", server.Config.Protocol, server.Config.Domain)
	}
	redirectTLS := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, dest+r.RequestURI, http.StatusMovedPermanently)
	}
	if err := http.ListenAndServe(origin, http.HandlerFunc(redirectTLS)); err != nil {
		log.Errorf("error ListenAndServe: %s", err)
		os.Exit(1)
	}
}

// Stop should not block. Return with a few seconds.
func (server Server) Stop(s service.Service) error {
	fmt.Println("clossing connection to main db")
	server.MainDB.Close()
	return nil
}

func (server *Server) load() error {
	if err := server.loadConfig(); err != nil {
		return err
	}
	//fmt.Printf("%#v\n", server.Config)
	//os.Exit(1)
	if err := server.connectDB(); err != nil {
		return err
	}
	server.configureHttpServer()
	return nil
}

func (server *Server) loadConfig() error {
	viper.SetEnvPrefix(config.EnvPrefix)
	viper.SetTypeByDefaultValue(true)
	server.registerEnvs(Config{})

	viper.SetDefault("REQUEST_TIMEOUT", 10)
	viper.SetDefault("DB_MAIN_MAX_IDLE_CONNS", 10)
	viper.SetDefault("DB_MAIN_MAX_OPEN_CONNS", 5)
	viper.SetDefault("DB_MAIN_CONN_MAX_IDLE_TIME", 300)
	viper.SetDefault("DB_MAIN_CONN_MAX_LIFETIME", 600)
	{ //PATH

		p := os.Getenv("GOPATH")
		if p == "" {
			p = build.Default.GOPATH
		}
		p = addSlash(p) + "src/" + config.SysPkgName + "/"
		viper.SetDefault("PATH", p)
	}

	err := viper.Unmarshal(&server.Config)
	if err != nil {
		return nil
	}
	return err
}
func (server *Server) registerEnvs(iv interface{}) {
	v := reflect.ValueOf(iv)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("mapstructure")
		tagValues := strings.Split(tag, ",")
		viper.BindEnv(tagValues[0])
	}
}

func (server *Server) connectDB() (err error) {
	fmt.Println("connecting to main db")
	if server.MainDB, err = sql.Open("postgres", string(server.Config.DBMainConnection)); err != nil {
		return err
	}
	server.MainDB.SetMaxIdleConns(server.Config.DBMainMaxIdleConns)
	server.MainDB.SetMaxOpenConns(server.Config.DBMainMaxOpenConns)
	server.MainDB.SetConnMaxIdleTime(time.Second * time.Duration(server.Config.DBMainConnMaxIdleTime))
	server.MainDB.SetConnMaxLifetime(time.Second * time.Duration(server.Config.DBMainConnMaxLifetime))
	return nil
}

func addSlash(s string) string {
	if s[len(s)-1] != '/' {
		return s + "/"
	}
	return s
}
