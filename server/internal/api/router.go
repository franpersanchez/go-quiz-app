package api

import (
	"go-quiz-app/server/internal/service"

	"github.com/gorilla/mux"
)

func Router(svc *service.QuizService) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/questions", GetQuestions(svc)).Methods("GET")
	r.HandleFunc("/submitAnswers", SubmitAnswers(svc)).Methods("POST")
	r.HandleFunc("/submissions", GetAllSubmissions(svc)).Methods("GET")
	return r
}
