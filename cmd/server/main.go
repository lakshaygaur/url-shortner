package main

import (
	log "github.com/sirupsen/logrus"

	"url_shortener/cmd/config"
	"url_shortener/internal/handler"
	"url_shortener/internal/service"
	"url_shortener/internal/storage"
	"url_shortener/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()
	utils.SetupLogger()

	if !cfg.DebugServer {
		log.Info("Running in production mode")
		gin.SetMode(gin.ReleaseMode)
	}

	// Setup API routes
	store := storage.NewMemoryStore()
	shortenService := service.NewShortenService(store)
	h := handler.NewHandler(shortenService)
	router := gin.Default()
	handler.RegisterRoutes(router, h)
	if err := router.Run(cfg.Host + ":" + cfg.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	log.Info("Server running on ", cfg.Host+":"+cfg.Port)
}
