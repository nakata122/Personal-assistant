package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
    "github.com/golang-jwt/jwt/v5"

	"server/internal/config"
	"server/internal/users"
	"server/internal/emails"
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
	
	//Database handling
	userData, err := GetUserData(token);
	if err != nil {
		log.Println(err);
		return;
    }

	var id int;
	curUser := users.GetUserByEmail(c, userData.Email);
	if curUser == nil {
		id = users.CreateUser(c, userData);
	} else {
		id = curUser.UserID;
	}

	myToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":    id,
        "email": userData.Email,
        "exp":   time.Now().Add(24 * time.Hour).Unix(),
    });

    session_token, err := myToken.SignedString([]byte(os.Getenv("JWT_SECRET")));
    if err != nil {
        log.Fatal(err);
    }

	setCookies(c, token, session_token);
	
	//Redirect to front-end

	c.Redirect(http.StatusFound, "http://localhost:" + os.Getenv("CLIENT_PORT") + "/dashboard");

	go func() {
		messages := GetMessages(c, token, id, 2);

		for _,message := range messages {
			var emailData emails.Email;
			emailData.UserID = id;
			emailData.Title = message.Title;
			emailData.Summary = message.Summary;
			emailData.Tags = []string{"formal", "urgent"};

			emails.CreateEmail(c, emailData);
		}
	}();

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

func setCookies(c *gin.Context, token *oauth2.Token, session_token string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "session_token",
		Value:    session_token,
		Expires:  time.Now().Add(24 * time.Hour),
        Path:     "/",
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	});

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



    