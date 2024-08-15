package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var ranking = &cobra.Command{
	Use:   "ranking",
	Short: "It serves all current submissions.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := showSubmissions(); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

var submissions []Result

func showSubmissions() error {
	resp, err := http.Get("http://localhost:8080/submissions")
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

	err = json.Unmarshal(body, &submissions)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	submissionsColor := color.New(color.FgYellow).Add(color.Bold)

	for _, submission := range submissions {
		submissionsColor.Printf("Submission -> Correct Answers: %d, Questions: %d, Score: %.2f\n", submission.Correct, submission.Total, submission.Score)
	}

	return nil
}
