package configtest

import (
	"database/sql"
	"log"

	"github.com/jtorz/phoenix-backend/app/config"
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
