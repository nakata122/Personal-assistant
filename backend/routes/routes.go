package routes

import (
	"github.com/gin-gonic/gin"
	
	"backend/controllers"
)

func RegisterPublicEndpoints(router *gin.Engine) {
	
	router.GET("./api/auth/google", controllers.GoogleLogin);
	router.GET("./api/auth/google_callback", controllers.GoogleCallback);
	router.GET("./api/auth/logout", controllers.Logout);
	router.GET("./api/ping", controllers.Ping);
}
