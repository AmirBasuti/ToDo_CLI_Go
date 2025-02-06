package cmd

import (
	"ToDo/internal/TodoBehavior"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// listCmd represents the list command.
var listCmd = &cobra.Command{
	Use:   "list", // Simplified Use message
	Short: "List all tasks.",
	Long: `List all tasks in your to-do list.

Example:
  ToDo list`, // More concise example
	Aliases: []string{"ls"}, // Use a shorter, common alias
	Args:    cobra.NoArgs,   // Explicitly state that no arguments are expected
	Run:     runList,        // Use a named function
}

func init() {
	rootCmd.AddCommand(listCmd)
}

// runList executes the list command logic.
func runList(cmd *cobra.Command, args []string) {

	err := TodoBehavior.List()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err) // Print errors to stderr
		os.Exit(1)                                 // Exit with an error code
	}
}
