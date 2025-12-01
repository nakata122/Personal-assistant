package controllers


import (
	"fmt"
	"time"
	"os"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	"backend/config"
)

func MyCookie(c *gin.Context) { 
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "test",
		Value:    "I am important cookie",
		Expires:  time.Now().Add(time.Duration(100) * time.Hour),
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	});
	c.Redirect(http.StatusFound, "http://localhost:5173/");
}

func Ping(c *gin.Context) { 
	c.JSON(200, gin.H {
		"message": "pong",
	});
}

func GoogleLogin(c *gin.Context) {
	// Redirect user to Google's OAuth consent screen
    url := config.Oauth.AuthCodeURL(os.Getenv("STATE"), oauth2.AccessTypeOffline);
	c.Redirect(http.StatusTemporaryRedirect, url);
}

func GoogleCallback(c *gin.Context) {
	//Check state
	state := c.Query("state");

    if state != os.Getenv("STATE") {
		fmt.Println("States don't Match!" + state);
		return;
    }

	//Get token coded
    code := c.Query("code");

	//Decode token
    token, err := config.Oauth.Exchange(c.Request.Context(), code);
    if err != nil {
		fmt.Println(err);
		return;
    }
	setCookies(c, token);

	//Get user info
    resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken);
    if err != nil {
		fmt.Println(err);
		return;
    }

	//Read user info
    userData, err := io.ReadAll(resp.Body);
    if err != nil {
		fmt.Println(err);
		return;
    }

	fmt.Println(string(userData));

	//Redirect to front-end
	c.Redirect(http.StatusFound, "http://localhost:5173/");
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