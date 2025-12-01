package main

import (
	"os"
	"github.com/gin-gonic/gin"

	"backend/config"
	"backend/routes"
)




func main() {
	config.LoadEnv();
	config.InitOauthConfig();

	router := gin.Default();
	routes.RegisterPublicEndpoints(router);

	port := os.Getenv("PORT");
	router.Run("localhost:" + port);
}