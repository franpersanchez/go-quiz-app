package service

import (
	"fmt"
	"go-quiz-app/server/internal/core"
	"go-quiz-app/server/internal/storage"
	models "go-quiz-app/server/pkg"
)

type QuizService struct {
	questionsStorage   *storage.QuestionsStorage
	submissionsStorage *storage.SubmissionsStorage
}

func NewQuizService(storage *storage.QuestionsStorage, submissionsStorage *storage.SubmissionsStorage) *QuizService {
	return &QuizService{
		questionsStorage:   storage,
		submissionsStorage: submissionsStorage}
}

func (s *QuizService) GetQuestions(amount *string, category *string, difficulty *string) []core.Question {
	return s.questionsStorage.GetQuestions(amount, category, difficulty)
}

func (s *QuizService) CheckAnswers(answers []models.Answer) models.Result {
	correctsCount := 0
	for _, answer := range answers {
		correctAnswer, exists := s.questionsStorage.GetAnswerResult(answer.QuestionID)
		if exists && correctAnswer == answer.Answer {
			correctsCount++
		}
	}
	score := float64(correctsCount) / float64(len(answers)) * 100
	ranking := s.CheckRanking(score)
	result := models.Result{
		Correct: correctsCount,
		Total:   len(answers),
		Score:   score,
		Ranking: ranking,
	}
	s.submissionsStorage.AddSubmissionResult(result)
	return result
}

func (s *QuizService) GetAllSubmissions() []models.Result {
	return s.submissionsStorage.GetAllSubmissionsResults()
}

func (s *QuizService) CheckRanking(score float64) string {
	submissions := s.submissionsStorage.GetAllSubmissionsResults()
	better := 0
	for _, submission := range submissions {
		if submission.Score < score {
			better++
		}
	}

	var ranking float64

	if len(submissions) > 0 {
		ranking = float64(better) / float64(len(submissions)) * 100
	} else {
		ranking = 0
	}
	return fmt.Sprintf("You were better than %.2f%% of all quizzers!", ranking)
}
