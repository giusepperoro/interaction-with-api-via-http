package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/giusepperoro/interaction-with-api-via-http/internal/config"
	"log"
)

const configFilePath = "./envs/config.yaml"

func main() {
	cfg, err := config.GetConfigFromFile(configFilePath)
	if err != nil {
		log.Fatal("unable to get config file name from env")
	}
	fmt.Println("cfg:", cfg)

	r := gin.Default()
	r.GET("./", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"cfg.SeverAddressUrl":    cfg.ServerAddressUrl,
			"cfg.PostgresConnectUrl": cfg.PostgresConnectUrl,
		})
	})
	r.Run(cfg.ServerAddressUrl)
}
