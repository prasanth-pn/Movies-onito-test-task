package main

import (
	"fmt"
	"log"

	"github.com/prasanth-pn/Movies-onito-test-task/pkg/config"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/db"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/di"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("the configuration files are not loaded %v", err)
	}
	fmt.Println(conf.DBSource)
	db.ConnectGorm(conf)
	server, dierr := di.InitializeEvent(conf)
	if dierr != nil {
		log.Fatal("canot start the server", dierr)
	} else {
		server.Start(conf)
	}

}
