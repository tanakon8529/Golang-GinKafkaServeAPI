package main

import (
	"fmt"
	"ginapi-gateway/controllers"
	"ginapi-gateway/docs"
	"ginapi-gateway/middleware"
	"ginapi-gateway/settings"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load environment variables
	config := settings.LoadEnv(".env")

	// Set up the Gin router
	router := gin.Default()

	// Set up CORS middleware options
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Set up your API routes
	router.POST(config.ApiPath+"/auth", controllers.AuthHandler())

	// For routes with middleware, you should use router.Group
	apiRoutes := router.Group(config.ApiPath)
	apiRoutes.Use(middleware.TokenAuthMiddleware())
	apiRoutes.GET("/health", controllers.HealthCheck)
	apiRoutes.POST("/kafka", controllers.SendKafkaMessage)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Gin Gateway API"
	docs.SwaggerInfo.Description = "This is a Headquarters Gateway"
	docs.SwaggerInfo.Version = config.ApiVersion
	docs.SwaggerInfo.Host = config.HostGateway
	docs.SwaggerInfo.BasePath = config.ApiPath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// Swagger endpoint
	log.Printf("Swagger endpoint: http://%s:%s/swagger/index.html", config.HostGateway, config.Port)
	swaggerURL := ginSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", config.HostGateway, config.Port))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	// Start the background jobs
	go middleware.StartJobs()

	// Run the server
	log.Printf("Server is running on %s:%s", config.Host, config.Port)
	router.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
}
