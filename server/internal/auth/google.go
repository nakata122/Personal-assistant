package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"

	"server/internal/emails"
	"server/internal/config"
	"server/internal/users"
)


func GetUserData(token *oauth2.Token) (*users.User, error) {
	//Get user info
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken);
	if err != nil {
		log.Println(err);
		return nil, err;
	}

	//Read user info
	decoder := json.NewDecoder(resp.Body);
	decoder.DisallowUnknownFields();

    var data users.User;
    decoder.Decode(&data);

	return &data, nil;
}

func GetMessageIDs(srv *gmail.Service, user string, max int64) []string {
	var ids []string;

	resp, err := srv.Users.Messages.List(user).MaxResults(max).Do();

	if err != nil {
		log.Printf("Unable to list messages: %v", err);
	}

	for _, m := range resp.Messages {
		ids = append(ids, m.Id);
	}

	return ids;
}


func GetMessages(c *gin.Context, token *oauth2.Token, user_id int, max int64) []emails.Email{
	client := config.Oauth.Client(c, token);

	srv, _ := gmail.NewService(c, option.WithHTTPClient(client));

	user := "me";
	ids := GetMessageIDs(srv, user, max);

	messages := []emails.Email{}

	for _, id := range ids {
		message, err := srv.Users.Messages.Get(user, id).Format("full").Do();
		if err != nil {
			log.Printf("Unable to get message: %v", err);
		}

		subject := ""
		for _, h := range message.Payload.Headers {
			if h.Name == "Subject" {
				subject = h.Value;
				break;
			}
		}
		

		var buffer bytes.Buffer;
		if(message.Payload.Body != nil && message.Payload.Body.Data != "") {
			data, _ := base64.URLEncoding.DecodeString(message.Payload.Body.Data);
			buffer.Write(data);
		}
		
		readBodyParts(message.Payload.Parts, &buffer);
		
		body := htmlToText(buffer.String());
		
		parsedEmail := emails.ParseEmail(c, body);
		messages = append(messages, emails.Email{UserID: user_id, Summary: parsedEmail, Title: subject});
	}

	return messages;
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

