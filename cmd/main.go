package main

import (
	"fmt"
	"log"

	"github.com/prasanth-pn/Movies-onito-test-task/pkg/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("the configuration files are not loaded %v", err)
	}
	fmt.Println(config.DBSource)

}
