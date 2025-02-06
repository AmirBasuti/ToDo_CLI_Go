package cmd

import (
	"ToDo/internal/TodoBehavior"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// addCmd represents the add command.
var addCmd = &cobra.Command{
	Use:   "add <task_description>", // Clarified Use message
	Short: "Add a new task to your to-do list.",
	Long: `Add a new task to your to-do list.  Provide the task description as an argument.

Example:
  ToDo add "Buy groceries"`, // Added a concrete example
	Aliases: []string{"a"},      // Use a shorter, more common alias
	Args:    cobra.ExactArgs(1), // Enforce exactly one argument
	Run:     runAdd,             // Use a named function
}

func init() {
	rootCmd.AddCommand(addCmd)

}

// runAdd executes the add command logic.
func runAdd(cmd *cobra.Command, args []string) {
	taskDescription := args[0]

	// Input validation (optional, but good practice):
	if len(taskDescription) == 0 {
		fmt.Fprintf(os.Stderr, "Error: Task description cannot be empty.\n")
		os.Exit(1)
		return // this is optional, since os.Exit will terminate
	}

	err := TodoBehavior.Add(taskDescription) // Call TodoBehavior.Add
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err) // Print errors to stderr
		os.Exit(1)                                 // Exit with an error code
	}
}
