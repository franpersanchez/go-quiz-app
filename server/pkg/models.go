package models

type Question struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
}

type Answer struct {
	QuestionID int    `json:"question_id"`
	Answer     string `json:"answer"`
}

type Result struct {
	Correct int     `json:"correct"`
	Total   int     `json:"total"`
	Score   float64 `json:"score"`
	Ranking string  `json:"ranking"`
}
