package cmd

import (
	"ToDo/internal/TodoBehavior"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// deleteCmd represents the delete command.
var deleteCmd = &cobra.Command{
	Use:   "delete <task_id>", // Improved Use message
	Short: "Delete a task by its ID.",
	Long: `Delete a task from your to-do list by providing its unique ID.

Example:
  ToDo delete 123`, // More concise example
	Aliases: []string{"del"},    // Removed redundant "delete" alias
	Args:    cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run:     runDelete,          // Use a named function for better readability
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Example of adding a flag (optional):  You could add a flag to force deletion without confirmation.
	// deleteCmd.Flags().BoolP("force", "f", false, "Force deletion without confirmation")
}

// runDelete executes the delete command logic.
func runDelete(cmd *cobra.Command, args []string) {
	taskIDStr := args[0]

	// Validate that the argument is a number *before* passing it to TodoBehavior.Delete.
	_, err := strconv.Atoi(taskIDStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Invalid task ID '%s'.  Task ID must be a number.\n", taskIDStr)
		os.Exit(1) // Exit with a non-zero code to indicate an error
		return     // this return is not really necessary
	}
	// Get the "force" flag value (if you added the flag).  This shows how to access flags.
	// force, _ := cmd.Flags().GetBool("force")

	// Add a confirmation prompt (optional, but good practice).
	if !confirmDeletion(taskIDStr) {
		fmt.Println("Deletion cancelled.")
		return
	}

	err = TodoBehavior.Delete(taskIDStr) // Now we can be sure we're passing a valid number string
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err) // Use Fprintf to print to stderr
		os.Exit(1)                                 // Exit with error code
	}
}

// confirmDeletion prompts the user for confirmation before deleting the task.
func confirmDeletion(taskID string) bool {
	fmt.Printf("Are you sure you want to delete task with ID %s? (y/N): ", taskID)
	var response string
	fmt.Scanln(&response) // Use Scanln to read a full line, including spaces.

	//  Handle various forms of "yes" (case-insensitive).
	switch response {
	case "y", "Y", "yes", "Yes", "YES":
		return true
	default:
		return false
	}
}
