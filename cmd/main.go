package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/giusepperoro/interaction-with-api-via-http/internal/config"
	"github.com/giusepperoro/interaction-with-api-via-http/internal/handlers"
	"log"
)

const configFilePath = "./envs/config.yaml"

func main() {
	cfg, err := config.GetConfigFromFile(configFilePath)
	if err != nil {
		log.Fatal("unable to get config file name from env")
	}
	fmt.Println("cfg:", cfg)
	c := handlers.NewAddToDatabaseHandler(cfg)
	router := gin.Default()
	router.GET("/", c.HandleAddToDatabase)
	router.Run(cfg.ServerAddressUrl)
}
