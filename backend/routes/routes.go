package routes

import (
	"github.com/gin-gonic/gin"
	
	"backend/controllers"
)

func RegisterPublicEndpoints(router *gin.Engine) {
	router.GET("./api/auth/google", controllers.GetGoogleAuth);
	router.GET("./api/auth/google_callback", controllers.GetGoogleCallback);
	router.GET("./api/ping", controllers.Ping);
}
