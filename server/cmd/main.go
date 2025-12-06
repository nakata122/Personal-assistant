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

    router.Static("/assets", "./dist/assets");

    router.NoRoute(func(c *gin.Context) {
        c.File("./dist/index.html");
    })

	http.InitMiddleware(router);

	http.InitRoutes(router);

	port := os.Getenv("PORT");
	router.Run("localhost:" + port);
}