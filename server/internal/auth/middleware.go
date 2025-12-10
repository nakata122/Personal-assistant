package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"server/internal/config"
)

// Middleware
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get token from HttpOnly cookie
        cookie, err := c.Cookie("session_token");
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing cookie"});
            return;
        }

        // Parse and verify JWT
        claims := jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		});
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"});
            return;
        }


        // check expiration
        if exp, ok := claims["exp"].(float64); ok {
            if int64(exp) < time.Now().Unix() {
                c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"});
                return;
            }
        }

        // Extract user info
        idFloat, okID := claims["id"].(float64);
        email, okEmail := claims["email"].(string);
        
        if !okID || !okEmail {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token data"});
            return;
        }

        user := config.ContextUser{
            ID:    int(idFloat),
            Email: email,
        }

        // Store user in context
        c.Set("user", user);
        c.Next();
    }
}