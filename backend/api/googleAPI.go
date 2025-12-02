package api

import (
	"backend/config"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func ReadUserData(token *oauth2.Token) (string, error) {
	//Get user info
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken);
	if err != nil {
		fmt.Println(err);
		return "", err;
	}

	//Read user info
	userData, err := io.ReadAll(resp.Body);
	if err != nil {
		return "", err;
	}

	return string(userData), nil;
}

func listMessageIDs(srv *gmail.Service, user string, max int64) []string {
    var ids []string;

    resp, err := srv.Users.Messages.List(user).MaxResults(max).Do();

    if err != nil {
        fmt.Printf("Unable to list messages: %v", err);
    }

    for _, m := range resp.Messages {
        ids = append(ids, m.Id);
    }

    return ids;
}

func ReadMessages(c *gin.Context, token *oauth2.Token) {
	client := config.Oauth.Client(c, token);

    srv, _ := gmail.NewService(c, option.WithHTTPClient(client));

    user := "me"
    ids := listMessageIDs(srv, user, 100);  // get first 100 message IDs
	
    message, err := srv.Users.Messages.Get(user, ids[0]).Format("full").Do();
	if err != nil {
		fmt.Printf("Unable to get message: %v", err);
	}

	body := "";
	if message.Payload.Body != nil && message.Payload.Body.Data != "" {
		data, _ := base64.URLEncoding.DecodeString(message.Payload.Body.Data);
		body = string(data)
	}

	fmt.Println(body);
}