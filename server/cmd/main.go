package main

import (
	"go-quiz-app/server/internal/api"
	"go-quiz-app/server/internal/service"
	"go-quiz-app/server/internal/storage"
	"log"
	"net/http"
)

func main() {
	//initialize in-memory storage
	questionsStorage, err := storage.InitializeStorage()

	if err != nil {
		log.Fatalf("Failed to initialize storage of Questions: %v", err)
	}

	submissionsStorage := storage.InitializeSubmissionsStorage()

	//initialize quiz Service
	quizService := service.NewQuizService(questionsStorage, submissionsStorage)

	//initialize the router for the API
	router := api.Router(quizService)

	log.Println("Backend server is starting on port: 8080")

	//starts the server using the defined router
	log.Fatal(http.ListenAndServe(":8080", router))
}
