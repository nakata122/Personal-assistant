package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"backend/config"
)

func loadEnv() {
	env, err := os.Open(".env");
	if err != nil {
		fmt.Println(err)
	}
	
	s := bufio.NewScanner(env);
	s.Split(bufio.ScanWords);

	for s.Scan() {
		key := s.Text();
		s.Scan();

		os.Setenv(key, s.Text());
	}

	defer env.Close();
}

func ping(c *gin.Context) { 
	c.JSON(200, gin.H {
		"message": "pong",
	})
}

func getGoogleAuth(c *gin.Context) {
	// Create random state to protect against CSRF
    state := "NODSINFsdiofniosf9f3j2rn039@$#@$(NQ)WR#@RN(N)";

	// Redirect user to Google's OAuth consent screen
    url := config.GetOauthConfig().AuthCodeURL(state, oauth2.AccessTypeOffline);
	c.Redirect(http.StatusMovedPermanently, url);

	fmt.Println(url);
}

func getGoogleCallback(c *gin.Context) {

	state := c.Query("state");

    if state != "NODSINFsdiofniosf9f3j2rn039@$#@$(NQ)WR#@RN(N)" {
		fmt.Println("States don't Match!" + state);
		return;
    }

    code := c.Query("code");


    token, err := config.GetOauthConfig().Exchange(c.Request.Context(), code);
    if err != nil {
		fmt.Println(err);
		return;
    }

    resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken);
    if err != nil {
		fmt.Println("User Data Fetch Failed");
		return;
    }

    userData, err := io.ReadAll(resp.Body);
    if err != nil {
		fmt.Println("JSON Data Fetch Failed");
		return;
    }

	fmt.Println(string(userData));


	c.Redirect(http.StatusMovedPermanently, "http://localhost:5173/");
}

func main() {
	loadEnv();

	router := gin.Default();
	router.GET("./api/auth/google", getGoogleAuth);
	router.GET("./api/auth/google_callback", getGoogleCallback);
	router.GET("./api/ping", ping);


	port := os.Getenv("PORT");
	fmt.Println("Starting server on port: " + port);
	router.Run("localhost:" + port);
}