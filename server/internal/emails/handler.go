package emails

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"server/internal/config"

	// "os/exec"

	"github.com/gin-gonic/gin"
)

func ParseEmail(c *gin.Context, message string) string{
	// cmd := exec.Command("python", "./scripts/script.py", message);

	// out, err := cmd.CombinedOutput();
	// if err != nil {
	// 	log.Println("Error:", err);
	// }

	// log.Println(string(out));

	payload := map[string]interface{}{
		"inputs": message,
		"parameters": map[string]interface{}{
			"max_length": 40, 
			"min_length": 0,
		},
	}

	jsonPayload, _ := json.Marshal(payload);

	apiToken := os.Getenv("HF_TOKEN");
	url := "https://router.huggingface.co/hf-inference/models/facebook/bart-large-cnn";
	// url := "https://router.huggingface.co/hf-inference/models/Falconsai/text_summarization";


	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload));
	req.Header.Set("Authorization", "Bearer "+apiToken);
	req.Header.Set("Content-Type", "application/json");

	type HFResponse []struct {
		SummaryText   string      `json:"summary_text"`
	}

	
	client := &http.Client{}
	res, err := client.Do(req);
	if err != nil {
		log.Println(err);
	}
	defer res.Body.Close();

	body, _ := io.ReadAll(res.Body);
	
	var result HFResponse;
	json.Unmarshal(body, &result);
	// log.Printf("Request result %v", result);

	if result != nil {
		return result[0].SummaryText;
	}
	return "";
}

func GetEmails(c *gin.Context) {
	val, ok := c.Get("user");
	if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"});
        return;
    }
	user, ok := val.(config.ContextUser) 
	if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"});
		return;
	}

	query := `SELECT * FROM emails WHERE user_id=$1`;
    rows, err := config.DbConn.Query(c, query, user.ID);
	if err != nil {
		log.Println(err);
		return;
	}
	defer rows.Close();

	emails := []Email{};

	for rows.Next() {
		var e Email;

		err := rows.Scan(&e.EmailID, &e.UserID, &e.Title, &e.Summary, &e.Score, &e.Tags);
		if err != nil {
			log.Println(err);
		}

		// log.Println(e);
		emails = append(emails, e);
	}

	c.JSON(200, emails);
}

func CreateEmail(c *gin.Context, emailData Email) {
    query := `INSERT INTO emails (user_id, title, summary, score, tags) VALUES ($1, $2, $3, $4, $5)`;
    _, err := config.DbConn.Exec(c, query, emailData.UserID, emailData.Title, emailData.Summary, emailData.Score, emailData.Tags);

    if err != nil {
		log.Println(err);
        return;
    }
}