/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ToDo/internal/TodoBehavior"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "ToDo",
	Short: "A CLI tool to manage your to-do list.",
	Long: `ToDo is a command-line interface (CLI) application built with Cobra 
that allows you to manage your to-do list efficiently.  You can add, 
complete, delete, and list tasks.`,
	Run: runRootCmd, // Use a separate function for the Run logic.
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err) // Consistent error output
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle") // Example flag, keep if needed.
}

// runRootCmd is the main function that handles user interaction for the root command.
func runRootCmd(cmd *cobra.Command, args []string) {
	var countune = true
	for countune {
		//clean the terminal
		fmt.Print("\033[H\033[2J")

		commands := cmd.Commands()

		// Display available commands to the user.
		fmt.Println("Available Commands:")
		for i, c := range commands {

			fmt.Printf("%d: %s - %s\n", i+1, c.Name(), c.Short)
		}

		// Get user's command choice.
		choice, err := getUserChoice(len(commands))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return // Exit the function on input error
		}

		// Execute the chosen command.  A switch statement is cleaner than if/else.
		switch choice {
		case 2: // Complete
			executeComplete()
		case 4: // Delete
			executeDelete()
		case 6: // List
			executeList()
		case 1: // Add
			executeAdd()
		default:
			fmt.Println("Invalid choice.") // Should not happen due to getUserChoice validation.
		}
		fmt.Println("Do you want to continue? (y/n)")
		var response string
		fmt.Scan(&response)
		if response == "n" {
			countune = false
		}
	}
}

// getUserChoice prompts the user for a command choice and validates the input.
func getUserChoice(maxChoice int) (int, error) {
	var choice int
	fmt.Print("Enter the number of the command: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %w", err) // Wrap the error
	}

	if choice < 1 || choice > maxChoice {
		return 0, fmt.Errorf("invalid choice: must be between 1 and %d", maxChoice)
	}
	return choice, nil
}

// --- Command Execution Functions ---
// These functions handle the specific actions for each command.

func executeComplete() {
	idStr, err := promptForInput("Enter the ID of the task to complete: ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	// Validate if idStr is a number
	if _, err := strconv.Atoi(idStr); err != nil {
		fmt.Fprintln(os.Stderr, "Error ID has to be a numeric value:", err)
		return
	}

	TodoBehavior.Complete(idStr)
}

func executeDelete() {
	idStr, err := promptForInput("Enter the ID of the task to delete: ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	// Validate if idStr is a number
	if _, err := strconv.Atoi(idStr); err != nil {
		fmt.Fprintln(os.Stderr, "Error ID has to be a numeric value:", err)
		return
	}
	TodoBehavior.Delete(idStr)
}

func executeList() {
	TodoBehavior.List()
}

func executeAdd() {
	task, err := promptForInput("Enter the task to add: ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	TodoBehavior.Add(task)
}

// promptForInput prompts the user for input and handles potential errors.
func promptForInput(prompt string) (string, error) {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", fmt.Errorf("input error: %w", err)
	}
	return input, nil
}
