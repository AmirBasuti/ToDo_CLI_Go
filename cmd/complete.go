package cmd

import (
	"ToDo/internal/TodoBehavior"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// completeCmd represents the complete command.
var completeCmd = &cobra.Command{
	Use:   "complete <task_id>", // Improved Use message
	Short: "Mark a task as complete.",
	Long: `Mark a task as complete by providing its unique ID.

Example:
  ToDo complete 42`, // More concise example
	Aliases: []string{"comp"},   // Removed redundant "complete" alias
	Args:    cobra.ExactArgs(1), // Ensure exactly one argument
	Run:     runComplete,        // Use a named function
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Example of adding a flag (optional): You could add a flag to unmark a task as complete.
	// completeCmd.Flags().BoolP("unmark", "u", false, "Unmark a task as complete (mark as incomplete)")
}

// runComplete executes the complete command logic.
func runComplete(cmd *cobra.Command, args []string) {
	taskIDStr := args[0]

	// Input validation: Check if the argument is a valid number.
	_, err := strconv.Atoi(taskIDStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Invalid task ID '%s'. Task ID must be a number.\n", taskIDStr)
		os.Exit(1) // Exit with a non-zero code on error
		return     // this return is not really needed
	}

	// Example of how to access a flag (if you added the "unmark" flag).
	// unmark, _ := cmd.Flags().GetBool("unmark")
	// if unmark {
	//     // Call a function to unmark the task (you'd need to implement this in TodoBehavior).
	//     err = TodoBehavior.UnmarkComplete(taskIDStr)
	// } else {
	//     err = TodoBehavior.Complete(taskIDStr)
	// }

	err = TodoBehavior.Complete(taskIDStr) // Now taskIDStr is guaranteed to be a string representation of a number
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err) // Print errors to stderr
		os.Exit(1)                                 // Exit with an error code
	}
}
