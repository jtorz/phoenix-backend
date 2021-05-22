package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jtorz/phoenix-backend/app/config"
	"github.com/jtorz/phoenix-backend/app/server"
	"github.com/kardianos/service"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	flag.Bool("install", false, "installs the server as a service.")
	flag.Bool("uninstall", false, "uninstalls the service.")
	flag.Bool("start", false, "starts the service.")
	flag.Bool("stop", false, "stops the service.")
	flag.Bool("status", false, "prints the service status.")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
}

func main() {
	v := viper.New()
	v.BindPFlags(pflag.CommandLine)
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
	if v.GetBool("install") {
		err = s.Install()
	} else if v.GetBool("uninstall") {
		err = s.Uninstall()
	} else if v.GetBool("start") {
		err = s.Start()
	} else if v.GetBool("stop") {
		err = s.Stop()
	} else if v.GetBool("status") {
		var st service.Status
		st, err = s.Status()
		printStatus(st)
	} else {
		err = s.Run()
	}
	if err != nil {
		log.Fatal(err)
	}
}

func printStatus(s service.Status) {
	switch s {
	case service.StatusRunning:
		fmt.Println("running")
	case service.StatusStopped:
		fmt.Println("stopped")
	default:
		fmt.Println("???")
	}
}
