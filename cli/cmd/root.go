package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "GO-QUIZ-APP is an interactive quiz game.",
	Long: `GO-QUIZ-APP is an interactive quiz game with a Command Line Interface (CLI) based on TRIVIA questions.
		Built by Fran Pérez Sánchez in Go.
		Complete documentation is available at https://github.com/franpersanchez/go-quiz-app`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.AddCommand(startQuiz)
	rootCmd.AddCommand(ranking)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
