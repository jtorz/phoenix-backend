package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/kardianos/service"

	"github.com/jtorz/phoenix-backend/app/config/serverconfig"
	"github.com/jtorz/phoenix-backend/utils/syslog"

	// postgres driver
	_ "github.com/lib/pq"
)

type Server struct {
	Config     *serverconfig.Config
	MainDB     *sql.DB
	HTTPServer *http.Server
	Redis      *redis.Pool
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
	err = server.configure()
	if err != nil {
		log.Errorf("can't load server config: %s", err)
		return err
	}

	if err = server.HTTPServer.ListenAndServe(); err != nil {
		return err
	}
	/* go server.runRedirect(log)
	log.Infof("listenAndServeTLS on port %d", server.Config.Port)
	if err := server.HTTPServer.ListenAndServeTLS(server.Config.Cert, server.Config.Key); err != nil && err != http.ErrServerClosed {
		log.Errorf("error ListenAndServeTLS: %s", err)
		os.Exit(1)
	} */
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

func (server *Server) configure() (err error) {
	server.Config, err = serverconfig.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading test config: %s", err)
	}

	server.MainDB, err = serverconfig.OpenPostgres(
		server.Config.DBMainConnection,
		server.Config.DBMainMaxOpenConns,
		server.Config.DBMainMaxIdleConns,
		server.Config.DBMainConnMaxLifetime,
		server.Config.DBMainConnMaxIdleTime,
	)
	if err != nil {
		return fmt.Errorf("connecting database: %w", err)
	}

	server.Redis, err = serverconfig.OpenRedis(
		server.Config.RedisAddress,
		server.Config.RedisPassword,
		server.Config.RedisMaxOpenConns,
		server.Config.RedisMaxIdleConns,
	)
	if err != nil {
		return fmt.Errorf("connecting redis: %w", err)
	}
	server.configureServices()
	return nil
}
