package main

import (
	"junior/internal/config"
	"junior/internal/handler"
	"junior/internal/logger"
	"junior/internal/model"
	"junior/internal/repository"
	"junior/internal/service"
	"log"

	_ "junior/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg := config.LoadConfig()
	logger.Init()

	db := repository.InitDB(*cfg)
	if err := db.AutoMigrate(&model.Subscription{}); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	repo := repository.NewSubscriptionRepository(db)
	svc := service.NewSubscriptionService(repo)
	h := handler.NewHandler(svc)

	r := gin.Default()
	h.RegisterRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(r.Run(":" + cfg.Server.Port))
}
