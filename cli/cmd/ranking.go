package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ranking = &cobra.Command{
	Use: "ranking",
	Run: func(cmd *cobra.Command, args []string) {
		if err := showQuestions(); err != nil {
			fmt.Println("Error:", err)
		}
	},
}
