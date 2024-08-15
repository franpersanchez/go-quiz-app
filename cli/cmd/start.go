package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

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

var questions []Question
var answers []Answer
var rating Result

var (
	amount     int
	difficulty string
	category   int
)

var startQuiz = &cobra.Command{
	Use:   "start",
	Short: "It provides a list of questions and possible answers. You can only choose 1 correct answer.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := showQuestions(); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	startQuiz.Flags().IntVarP(&amount, "amount", "a", 10, "Number of questions to fetch")
	startQuiz.Flags().IntVarP(&category, "category", "c", 9, "Category ID")
	startQuiz.Flags().StringVarP(&difficulty, "difficulty", "d", "easy", "Difficulty level (easy, medium, hard)")
}

func showQuestions() error {
	url := fmt.Sprintf("http://localhost:8080/questions?amount=%d&difficulty=%s&category=%d", amount, difficulty, category)
	println(url)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching questions: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response code")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	err = json.Unmarshal(body, &questions)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	// Define color styles
	titleColor := color.New(color.BgYellow).Add(color.Bold)
	secondTitleColor := color.New(color.FgYellow)
	optionColor := color.New(color.FgGreen)
	errorColor := color.New(color.FgRed)

	secondTitleColor.Println("///////////////////////////")
	titleColor.Println("[ WELCOME TO THE QUIZ APP ]")
	secondTitleColor.Println("///////////////////////////")
	secondTitleColor.Println("You will be given a set of questions. Please, choose the correct answer to aim for the best results!")
	secondTitleColor.Println("///////////////////////////")

	for _, q := range questions {
		decodedQuestion := fmt.Sprintf("Question %d: %s", q.ID+1, html.UnescapeString(q.Question))
		prompt := promptui.Select{
			Label: decodedQuestion,
			Items: q.Options,
			Templates: &promptui.SelectTemplates{
				Active:   fmt.Sprintf("%s {{ . | %s }}", promptui.IconSelect, "cyan"),
				Inactive: "{{ . }}",
				Selected: fmt.Sprintf(`%s {{ . | %s }}`, promptui.IconGood, "cyan"),
			},
			HideHelp: true,
		}

		_, result, err := prompt.Run()
		if err != nil {
			errorColor.Printf("An error occurred while selecting an option: %v\n", err)
			return fmt.Errorf("an error occurred while selecting an option: %w", err)
		}

		answers = append(answers, Answer{
			QuestionID: q.ID,
			Answer:     result,
		})
	}

	if len(answers) == amount {
		secondTitleColor.Println("Quiz Completed! Here are your answers:")

		for _, ans := range answers {
			optionColor.Printf("Question %d: %s\n", ans.QuestionID+1, ans.Answer)
		}

		rating, err := checkRating(answers)
		if err != nil {
			return err
		}

		secondTitleColor.Printf("Result: Correct %d / %d\n", rating.Correct, rating.Total)
		secondTitleColor.Printf("Score: %.2f\n", rating.Score)
		secondTitleColor.Printf("Ranking: %s\n", rating.Ranking)

	}

	return nil
}

func checkRating(answers []Answer) (Result, error) {
	marshalledAnswers, err := json.Marshal(answers)
	if err != nil {
		return Result{}, fmt.Errorf("impossible to marshall teacher: %s", err)
	}
	resp, err := http.Post("http://localhost:8080/submitAnswers", "application/json", bytes.NewBuffer(marshalledAnswers))
	if err != nil {
		return Result{}, fmt.Errorf("error fetching questions: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Result{}, fmt.Errorf("received non-200 response code")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{}, fmt.Errorf("error reading response: %w", err)
	}

	err = json.Unmarshal(body, &rating)
	if err != nil {
		return Result{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return rating, nil

}
