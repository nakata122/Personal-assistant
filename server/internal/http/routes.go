package http

import (
	"github.com/gin-gonic/gin"
	
	"server/internal/auth"
)

func InitRoutes(router *gin.Engine) {
	
	router.GET("./api/auth/google", auth.GoogleLogin);
	router.GET("./api/auth/google_callback", auth.GoogleCallback);
	router.GET("./api/auth/logout", auth.Logout);
	router.GET("./api/ping", auth.Ping);
}
