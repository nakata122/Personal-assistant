package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	"server/internal/config"
	"server/internal/user"
)

func Ping(c *gin.Context) { 
	log.Println(c.Request.Cookies());

	c.JSON(200, gin.H{"Message": "pong"});
}

func GoogleLogin(c *gin.Context) {
	// Redirect user to Google's OAuth consent screen
    url := config.Oauth.AuthCodeURL(os.Getenv("STATE"));
	c.Redirect(http.StatusTemporaryRedirect, url);
}

func GoogleCallback(c *gin.Context) {
	//Check state
	state := c.Query("state");
    if state != os.Getenv("STATE") {
		log.Println("States don't Match!" + state);
		return;
    }

	//Get token coded
    code := c.Query("code");

	//Decode token
    token, err := config.Oauth.Exchange(c, code);
    if err != nil {
		log.Println(err);
		return;
    }
	setCookies(c, token);
	
	userData, err := GetUserData(token);
	if err != nil {
		log.Println(err);
		return;
    }

	user.CreateUser(c, userData);

	GetMessages(c, token, 10);

	//Redirect to front-end
	c.Redirect(http.StatusFound, "http://localhost:5173/dashboard");
}

func Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
        Path:     "/",
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
}

func setCookies(c *gin.Context, token *oauth2.Token) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    token.AccessToken,
		Expires:  token.Expiry,
        Path:     "/",
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	});

	// NOTE: Refresh token is only issued at the first consent
	if token.RefreshToken != "" {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "refresh_token",
			Value:    token.RefreshToken,
			Expires:  time.Now().Add(time.Duration(100) * time.Hour),
        	Path:     "/",
			HttpOnly: false,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		});
	}
}



    