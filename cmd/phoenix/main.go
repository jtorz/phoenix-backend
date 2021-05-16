package main

import (
	"flag"
	"log"

	"github.com/jtorz/phoenix-backend/app/config"
	"github.com/jtorz/phoenix-backend/app/server"
	"github.com/kardianos/service"
)

func main() {
	install := flag.Bool("install", false, "instala el servicio")
	uninstall := flag.Bool("uninstall", false, "desinstala el servicio")
	flag.Parse()
	svcConfig := &service.Config{
		Name:        config.SysKey + "-server",
		DisplayName: config.SysKey + "-server",
		Description: config.SysName + " backend server",
	}

	prg := server.NewServer()
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	if *install {
		err = s.Install()
	} else if *uninstall {
		err = s.Uninstall()
	} else {
		err = s.Run()
	}
	if err != nil {
		serviceLogger, errLog := s.Logger(nil)
		if errLog != nil {
			log.Fatal(errLog)
		}
		serviceLogger.Error(err)
	}
}
