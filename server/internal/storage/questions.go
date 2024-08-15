package storage

import (
	"encoding/json"
	"fmt"
	"go-quiz-app/server/internal/core"
	"net/http"
)

type QuestionsStorage struct {
	questions []core.Question
}

type TriviaApiResponse struct {
	ResponseCode int            `json:"response_code"`
	Results      []TriviaResult `json:"results"`
}

type TriviaResult struct {
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Category         string   `json:"category"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

// InitializeStorage initializes the storage and fetches new questions
func InitializeStorage() (*QuestionsStorage, error) {
	storage := &QuestionsStorage{
		questions: []core.Question{},
	}
	return storage, nil
}

// fetchNewQuestions fetches questions from Trivia API and stores them
func (s *QuestionsStorage) fetchNewQuestions(amount *string, category *string, difficulty *string) error {

	apiURL := fmt.Sprintf("https://opentdb.com/api.php?amount=%s&category=%s&difficulty=%s&type=multiple", *amount, *category, *difficulty)
	response, err := http.Get(apiURL)
	if err != nil {
		return fmt.Errorf("error fetching questions: %w", err)
	}
	defer response.Body.Close()

	var apiResponse TriviaApiResponse
	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		return fmt.Errorf("error decoding JSON response from Trivia API: %w", err)
	}

	var questionsGathered []core.Question
	for index, result := range apiResponse.Results {
		options := append(result.IncorrectAnswers, result.CorrectAnswer)
		questionsGathered = append(questionsGathered, core.Question{
			ID:             index,
			Question:       result.Question,
			Options:        options,
			Correct_answer: result.CorrectAnswer,
		})
	}

	s.questions = questionsGathered
	return nil
}

func (s *QuestionsStorage) GetAnswerResult(questionId int) (string, bool) {
	for _, question := range s.questions {
		if question.ID == questionId {
			return question.Correct_answer, true
		}
	}
	return "", false
}

func (s *QuestionsStorage) GetQuestions(amount *string, category *string, difficulty *string) []core.Question {
	s.fetchNewQuestions(amount, category, difficulty)
	return s.questions
}
