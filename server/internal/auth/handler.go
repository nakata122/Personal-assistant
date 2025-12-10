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

	curUser := users.GetUserByEmail(c, userData.Email);
	if curUser == nil {
		users.CreateUser(c, curUser);
	}

	session_token := CreateJWTToken(curUser);
	setCookies(c, token, session_token);
	
	//Redirect to front-end

	if(os.Getenv("ENV") == "PRODUCTION") {
		c.Redirect(http.StatusFound, os.Getenv("URL") + "/dashboard");
	} else {
		c.Redirect(http.StatusFound, "http://localhost:5173/dashboard");
	}

	go func() {
		messages := GetMessages(c, token, curUser.UserID, 2);

		for _,message := range messages {
			message.Tags = []string{"formal", "urgent"};

			emails.CreateEmail(c, message);
		}
	}();
}

func RegisterGuest(c *gin.Context) {
	var userData users.User;
	userData.GoogleID = "";
	userData.Role = users.RoleGuest;
	users.CreateUser(c, &userData);

	// updatedData := users.GetUserByID(c, id);
	session_token := CreateJWTToken(&userData);
	setCookies(c, nil, session_token);

	c.JSON(200, gin.H{"Message": "Guest user succesfully created"});
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

	if token != nil {
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

	
}

func CreateJWTToken(userData *users.User) string{
	myToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":    userData.UserID,
        "email": userData.Email,
        "exp":   time.Now().Add(24 * time.Hour).Unix(),
    });

    session_token, err := myToken.SignedString([]byte(os.Getenv("JWT_SECRET")));
    if err != nil {
        log.Println(err);
		return "";
    }

	return session_token;
}

    