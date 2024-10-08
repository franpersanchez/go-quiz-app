package api

import (
	"encoding/json"
	"fmt"
	"go-quiz-app/server/internal/service"
	models "go-quiz-app/server/pkg"
	"net/http"
)

func GetQuestions(svc *service.QuizService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		amount := query.Get("amount")
		category := query.Get("category")
		difficulty := query.Get("difficulty")
		questions, err := svc.GetQuestions(&amount, &category, &difficulty)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching questions from storage: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(questions)
	}
}

func SubmitAnswers(svc *service.QuizService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var answers []models.Answer
		if err := json.NewDecoder(r.Body).Decode(&answers); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result := svc.CheckAnswers(answers)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func GetAllSubmissions(svc *service.QuizService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		submissions := svc.GetAllSubmissions()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(submissions)
	}
}
