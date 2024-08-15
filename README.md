# Quiz App

This project is a very simple CLI quiz application built with Golang. 
The project consists of a backend server that provides quiz questions and processes submissions, along with a CLI client that enables users to customize quiz parameters, initiate the quiz, and view total submissions.

---

## User stories/Use cases: 

- User should be able to get questions with a number of answers
- User should be able to select just one answer per question.
- User should be able to answer all the questions and then post his/hers answers and get back how many correct answers they had, displayed to the user.
- User should see how well they compared to others that have taken the quiz, eg. "You were better than 60% of all quizzers"


## Features

- **Backend Server**: Serves quiz questions, submissions and handles answers.
    
    - /questions (GET) : it serves a certain number of questions and answers
    - /submitAnswers: it accepts a list of questions(id) and the selected answer.
    - /submissions: it serves a list of previous submissions, showing the score.

- **CLI Client (customizable Quiz)**: Allows users to take the quiz directly from the command line. Users can select the amount of questions, the category, and the difficulty of the quiz.

## Technology Used

- **Golang:** Backend and CLI implementation.
- **Cobra:** CLI library for creating powerful modern CLI applications.
- **promptui** CLI library providing a simple interface to create command-line prompts.

## Project Structure

```
go-quiz-app/
├── cli/
│   ├── main.go
│   └── cmd/
│       ├── root.go
│       └── start.go
│       └── ranking.go
│   ├── go.mod
│   ├── go.sum
├── server/
│   ├── main.go
│   ├── cmd/
│   ├── internal/
│       ├── api/
│       │   ├── handlers.go
│       │   └── router.go
│       ├── core/
│       │   ├── quiz.go
│       ├── service/
│       │   ├── quiz_service.go
│       ├── storage/
│       │   ├── questions.go
│       │   └── ranking.go
│   ├── pkg/
│   │   └── models.go
│   ├── go.mod
│   ├── go.sum
├── .gitignore
└── README.md
```