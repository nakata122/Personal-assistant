package emails


type Email struct {
	EmailID  int      `json:"email_id"`
	UserID   int      `json:"user_id"`
	Title    string   `json:"title"`
	Summary  string   `json:"summary"`
	Score    float32  `json:"score"`
	Tags     []string `json:"tags"`
}