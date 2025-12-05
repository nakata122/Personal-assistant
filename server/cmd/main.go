package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"server/internal/config"
	"server/internal/http"
)




func main() {
	config.LoadEnv();
	config.InitOauthConfig();
	config.ConnectDb();

	router := gin.Default();

	http.InitMiddleware(router);

	http.InitRoutes(router);

	port := os.Getenv("PORT");
	router.Run("localhost:" + port);
}