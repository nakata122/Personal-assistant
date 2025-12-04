package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/config"
	"backend/routes"
)




func main() {
	config.LoadEnv();
	config.InitOauthConfig();
	config.ConnectDb();

	router := gin.Default();

	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
	}));
	// router.Use(cors.Default());

	routes.RegisterPublicEndpoints(router);

	port := os.Getenv("PORT");
	router.Run("localhost:" + port);
}