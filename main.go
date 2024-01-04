package main

import (
	"errors"
	"flag"
	"fmt"
	"iam-server/config"
	"iam-server/db"
	"iam-server/logger"
	"iam-server/router"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var (
	env   = flag.String("environment", "dev", "Environment for Config")
	debug = flag.Bool("debug", true, "Debug Mode")
)

func main() {
	flag.Parse()
	if *env == "" {
		panic("No Environment Defined")
	}

	c := make(chan error)

	// Logging
	go logger.InitLogger(*debug)
	go startServer(c)

	err := <-c
	if err != nil {
		panic(c)
	}
}

func startServer(c chan error) {
	defer close(c)

	// Read Config
	err := config.ReadConfig(*env)
	if err != nil {
		fmt.Println(err.Error())
		c <- err
	}

	// Initialize Database
	err = db.ConnectMariadb()
	if err != nil {
		fmt.Println(err.Error())
		c <- err
	}

	// Get Server Port
	port := viper.Get("port")
	if port == nil || port == "" {
		c <- errors.New("port undefined")
	}

	app := echo.New()
	router.RegisterRoutes(app)
	app.Start(fmt.Sprintf(":%d", port))
}
