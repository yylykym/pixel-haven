package server

import (
	"github.com/gin-gonic/gin"
	"os"
	"pixel-haven/interval/config"
	"pixel-haven/interval/event"
)

func Start() {
	config.Init()
	go event.CreateTopic("upload")
	go event.StartConsumer()
	engine := gin.New()
	registerRoutes(engine)
	err := engine.Run(":8080")
	if err != nil {
		os.Exit(1)
		return
	}
}
