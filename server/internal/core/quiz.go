package core

type Question struct {
	ID             int      `json:"id"`
	Question       string   `json:"question"`
	Options        []string `json:"options"`
	Correct_answer string   `json:"correct_answer"`
	Answer         string   `json:"-"`
}

type Quiz struct {
	Questions []Question
}
