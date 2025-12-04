package controllers

import (
	"backend/config"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	// "io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"

	"backend/models"
)


func ReadUserData(token *oauth2.Token) (*models.User, error) {
	//Get user info
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken);
	if err != nil {
		fmt.Println(err);
		return nil, err;
	}

	//Read user info
	decoder := json.NewDecoder(resp.Body);
	decoder.DisallowUnknownFields();

    var data models.User;
    decoder.Decode(&data);

	return &data, nil;
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

func extractText(n *html.Node) string {
    if n.Type == html.TextNode {
        return n.Data;
    }
    var sb strings.Builder
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        sb.WriteString(extractText(c));
    }
    return sb.String();
}

func htmlToText(htmlStr string) string {
    doc, err := html.Parse(strings.NewReader(htmlStr));
    if err != nil {
        return htmlStr;
    }
    text := extractText(doc);
    return strings.Join(strings.Fields(text), " ");
}

func readBodyParts(parts []*gmail.MessagePart, text *bytes.Buffer) {

	for _, part := range parts {

		if(part.Body != nil && part.Body.Data != "") {
			data, _ := base64.URLEncoding.DecodeString(part.Body.Data);
			(*text).Write(data);
		}

		if len(part.Parts) > 0 {
            readBodyParts(part.Parts, text);
        }
	}
}

func ReadMessages(c *gin.Context, token *oauth2.Token) {
	client := config.Oauth.Client(c, token)

	srv, _ := gmail.NewService(c, option.WithHTTPClient(client))

	user := "me"
	ids := listMessageIDs(srv, user, 100) // get first 100 message IDs

	message, err := srv.Users.Messages.Get(user, ids[0]).Format("full").Do();
	if err != nil {
		fmt.Printf("Unable to get message: %v", err);
	}

    var buffer bytes.Buffer
	if(message.Payload.Body != nil && message.Payload.Body.Data != "") {
		data, _ := base64.URLEncoding.DecodeString(message.Payload.Body.Data);
		buffer.Write(data);
	}

	readBodyParts(message.Payload.Parts, &buffer);

	body := htmlToText(buffer.String());
	
	ExtractInfo(c, body);
}
