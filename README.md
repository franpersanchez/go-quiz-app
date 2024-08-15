# Quiz App

This project is a very simple CLI quiz application built with Golang. 
The project consists of a backend server that provides quiz questions and processes submissions, along with a CLI client that enables users to customize quiz parameters, initiate the quiz, and view the rest of submissions.

---

## About The Project
### User stories/Use cases: 

- User should be able to get questions with a number of answers
- User should be able to select just one answer per question.
- User should be able to answer all the questions and then post his/hers answers and get back how many correct answers they had, displayed to the user.
- User should see how well they compared to others that have taken the quiz, eg. "You were better than 60% of all quizzers"


### Features

- **Backend Server**: The backend server is responsible for serving quiz questions and processing answers. It utilizes resources from the [Trivia API](https://opentdb.com/api_config.php) to provide unique, engaging, and varied questions each time the quiz is initiated.
    
    - /questions (GET) : Retrieves a specified number of quiz questions and their possible answers.
    - /submitAnswers(POST): Accepts a list of question IDs and the selected answers. It processes the submission and sends a response with the score and ranking position.
    - /submissions(GET): Returns a list of previous submissions, including scores and details of each quiz attempt.

- **CLI Client (customizable Quiz)**: The command-line interface (CLI) client offers users an interactive and customizable quiz experience. Users can easily start the quiz directly from the command line and tailor the quiz to their preferences (optional) by selecting the number of questions, category, and difficulty level.

**Customizable Parameters**:

- **Amount of Questions**: Choose the number of questions for the quiz.
- **Category**: Select from a variety of categories to focus on specific topics of interest.
- **Difficulty**: Adjust the quiz difficulty level (easy, medium, or hard) to match your knowledge and challenge level.

### Technology Used

- **Golang:** Backend and CLI implementation.
- **Cobra:** CLI library for creating powerful modern CLI applications.
- **promptui** CLI library providing a simple interface to create 
command-line prompts.
- **TRIVIA API**: An external API used to source quiz questions. The Trivia API provides a free vast database of questions across various categories and difficulty levels, ensuring that each quiz session is unique and engaging for the user.

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


## Getting Started

### Prerequisites
Before you begin, ensure you have the following installed on your system:

- **Go**: Ensure you have Go installed. You can download it from the [official website](https://go.dev/dl/).
- **Git**: Ensure you have Git installed for cloning the repository. You can download it from the official website.

### Installation
Follow these steps to set up and run the project locally:

1. **Clone this Repository**:

    ```
    git clone https://github.com/franpersanchez/go-quiz-app

    cd go-quiz-app
    ```
2. **Install Dependencies**:

    Navigate to the server directory and install the necessary dependencies:

    ```
    cd server

    go mod tidy
    ```
    Then, navigate to the cli directory and install the necessary dependencies:

    ```
    cd cli

    go mod tidy
    ```
## Usage
The CLI client allows you to start the quiz and customize it by specifying the number of questions, category, and difficulty level. After completing the quiz, your results will be displayed along with a comparison to other participants.
You can choose to play the Quiz (command **start**) or to consult the current submissions (command **ranking**)

1. **Running the Backend Server**:

     Navigate back to the server directory and start the backend server:

     ```
     go run main.go
     ```
     The server will start locally on port 8080 by default.

2. **Using the CLI Client**:

    Open a **new terminal**, navigate to the cli directory, and run the CLI client:

    ```
    cd cli
    go run main.go start
    ```
    By default, you will be given a set of questions with the default values of amount:10, category: 9 (general knowledge) and difficulty: easy. 

    If you prefer to customize your quiz you can use the flags

    ```
    go run main.go start -a <AMOUNT> -c <CATEGORY> -d <DIFFICULTY>
    ```
    Replace < AMOUNT >, < DIFFICULTY >, and < CATEGORY > with your desired values. 
    

    For example:
    ```
    go run main.go start -a 5 -c 20 -d "hard"
    ```
    or
     ```
    go run . start -a 20 -c 10 -d "medium"
    ```


## Example Commands
### Quiz

Start a quiz with 10 questions of medium difficulty in the general knowledge category:

```
go run main.go start -a 10 -c 9 -d "medium"
```` 
For indicating a certain category you can reference this table:
| Number | Category                |
|--------|-------------------------|
| 9      | General knowledge       |
| 10     | Books                   |
| 11     | Films                   |
| 12     | Music                   |
| 13     | Musicals & Theaters     |
| 14     | TV                      |
| 15     | Video Games             |
| 16     | Board Games             |
| 17     | Science & Nature        |
| 18     | Computers               |
| 19     | Mathematics             |
| 20     | Mythology               |
| 21     | Sports                  |
| 22     | Geography               |
| 23     | History                 |
| 24     | Politics                |
| 25     | Art                     |
| 26     | Celebrities             |
| 27     | Animals                 |


### Ranking
View the ranking of all submissions:

```
go run main.go ranking
```

you can get more information about the commands using the "--help" command:

```
go run . start --help
```

## Resources
- [golang docs](https://go.dev/doc/)
- [cobra](https://github.com/spf13/cobra)
- [promptui](https://github.com/manifoldco/promptui)
- [Trivia API](https://opentdb.com/api_config.php)