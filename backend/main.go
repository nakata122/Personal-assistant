package main

import (
	"os"
	"time"
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"

	"backend/config"
	"backend/routes"
)




func main() {
	config.LoadEnv();
	config.InitOauthConfig();

	router := gin.Default();

	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	routes.RegisterPublicEndpoints(router);

	port := os.Getenv("PORT");
	router.Run("localhost:" + port);
}