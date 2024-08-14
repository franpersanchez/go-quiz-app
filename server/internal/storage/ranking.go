package storage

import (
	models "go-quiz-app/server/pkg"
	"sync"
)

type SubmissionsStorage struct {
	mu      sync.Mutex
	results []models.Result
}

func InitializeSubmissionsStorage() *SubmissionsStorage {
	return &SubmissionsStorage{}
}

func (rs *SubmissionsStorage) AddSubmissionResult(result models.Result) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	rs.results = append(rs.results, result)
}

func (rs *SubmissionsStorage) GetAllSubmissionsResults() []models.Result {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	return rs.results
}
