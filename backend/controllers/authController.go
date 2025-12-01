package controllers


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"net/http"
	"golang.org/x/oauth2"

	"backend/config"
)


func Ping(c *gin.Context) { 
	c.JSON(200, gin.H {
		"message": "pong",
	})
}

func GetGoogleAuth(c *gin.Context) {
	// Redirect user to Google's OAuth consent screen
    url := config.Oauth.AuthCodeURL(os.Getenv("STATE"), oauth2.AccessTypeOffline);
	c.Redirect(http.StatusMovedPermanently, url);
}

func GetGoogleCallback(c *gin.Context) {
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
	c.Redirect(http.StatusMovedPermanently, "http://localhost:5173/");
}