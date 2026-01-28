// Package main is the entry point of the application
package main

import (
	"fmt"
	"main-service/internal/database"
	"main-service/internal/link"
	"main-service/internal/metrics"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	_ "main-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Swagger PseudoLink API
// @version 1.0
// @description Generated API-cellers on server

// @contact.url https://github.com/achi3v3
// @contact.email aamir-tutaev@mail.ru

// @host localhost:8080
// @BasePath /
// func main() {
// 	logger := logrus.New()
// 	logger.SetLevel(logrus.DebugLevel)
// 	logger.Info("[ START ]")
// 	redisClient := database.Init(logger)
// 	defer func() {
// 		if err := redisClient.Close(); err != nil {
// 			logger.Warnf("TestCreate: error with close %v", err)
// 		}
// 	}()

// 	linkRepo := link.NewRepository(redisClient)
// 	linkService := link.NewService(linkRepo)
// 	linkHandler := link.NewHandler(linkService)

// 	router := gin.Default()
// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
// 	linkGroup := router.Group("/link")
// 	{
// 		linkGroup.POST("/create", linkHandler.Create)
// 		linkGroup.GET("/get", linkHandler.GetPseudo)
// 		linkGroup.DELETE("/delete", linkHandler.Delete)
// 	}
// 	router.GET("/:shortID", linkHandler.Redirect)
// 	router.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Привет, Gin!")
// 	})

// 	go func() {
// 		if err := metrics.Listen("127.0.0.1:9090"); err != nil {
// 			logger.Warnf("error with listen metrics %v", err)
// 		}
// 	}()
// 	logger.Infof("Run server :8080")
// 	if err := router.Run(":8080"); err != nil {
// 		logger.Fatalf("Failed to run server: %v", err)
// 	}
// 	logger.Infof("[ STOP ]")
// }

func main() {
	// ДОБАВЬ ДЕБАГ
	fmt.Fprintf(os.Stderr, "=== MAIN START: %v ===\n", time.Now())

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.Info("[ START ]")

	fmt.Fprintln(os.Stderr, "STEP 1: Logger created")

	redisClient := database.Init(logger)
	fmt.Fprintln(os.Stderr, "STEP 2: Redis initialized")

	defer func() {
		if err := redisClient.Close(); err != nil {
			logger.Warnf("TestCreate: error with close %v", err)
		}
		fmt.Fprintln(os.Stderr, "STEP 9: Redis closed")
	}()

	linkRepo := link.NewRepository(redisClient)
	fmt.Fprintln(os.Stderr, "STEP 3: Repository created")

	linkService := link.NewService(linkRepo)
	fmt.Fprintln(os.Stderr, "STEP 4: Service created")

	linkHandler := link.NewHandler(linkService)
	fmt.Fprintln(os.Stderr, "STEP 5: Handler created")

	router := gin.Default()
	fmt.Fprintln(os.Stderr, "STEP 6: Router created")
	if os.Getenv("CI") != "true" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	linkGroup := router.Group("/link")
	{
		linkGroup.POST("/create", linkHandler.Create)
		linkGroup.GET("/get", linkHandler.GetPseudo)
		linkGroup.DELETE("/delete", linkHandler.Delete)
	}
	router.GET("/:shortID", linkHandler.Redirect)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Привет, Gin!")
	})

	fmt.Fprintln(os.Stderr, "STEP 7: Routes configured")

	go func() {
		fmt.Fprintln(os.Stderr, "STEP 8: Starting metrics server...")
		if err := metrics.Listen("127.0.0.1:9090"); err != nil {
			logger.Warnf("error with listen metrics %v", err)
		}
		fmt.Fprintln(os.Stderr, "STEP 8a: Metrics server stopped")
	}()

	logger.Infof("Run server :8080")
	fmt.Fprintln(os.Stderr, "STEP 9: Starting Gin server on :8080")

	if err := router.Run(":8080"); err != nil {
		logger.Fatalf("Failed to run server: %v", err)
		fmt.Fprintln(os.Stderr, "STEP 10: Server failed")
	}

	logger.Infof("[ STOP ]")
	fmt.Fprintln(os.Stderr, "=== MAIN END ===")
}
