package emails


type Email struct {
	EmailID    int      `json:"email_id"`
	UserID     int      `json:"user_id"`
	Title      string   `json:"title"`
	Summary    string   `json:"summary"`
	ProfilePic string   `json:"profilepic"`
	Score      float32  `json:"score"`
	Tags       []string `json:"tags"`
}