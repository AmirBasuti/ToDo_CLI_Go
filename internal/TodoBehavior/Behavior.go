package TodoBehavior

import (
	"ToDo/internal/Database"
	"ToDo/internal/Model"
	"fmt"
	"strconv"
	"strings"
)

// Constants for table formatting.  Makes it easy to adjust the look.
const (
	idHeader        = "ID"
	taskHeader      = "Task"
	stateHeader     = "State"
	completeState   = "Complete"
	incompleteState = "Incomplete"
	separatorWidth  = 4 //  "| " + " |"  = 4 characters

)

// Helper function to check if the database connection is initialized.
func checkDB() error {
	if Database.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	return nil
}

// Add creates a new task with the given task text.
// It returns an error if the task creation fails, or if the database is not connected.
func Add(taskTxt string) error {
	if err := checkDB(); err != nil {
		return err
	}

	db := Database.DB
	task := Model.Todo{Task: taskTxt}
	if res := db.Create(&task); res.Error != nil {
		return fmt.Errorf("create task failed: %w", res.Error) // Use %w for error wrapping
	}
	fmt.Printf("Your task was successfully created with ID: %d\n\n", task.ID)
	return nil
}

// Complete marks a task as complete given its ID.
// It returns an error if the task update fails, or if the database is not connected.
func Complete(arg string) error {
	id, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	if err := checkDB(); err != nil {
		return err
	}
	db := Database.DB

	res := db.Model(&Model.Todo{}).Where("id = ?", uint(id)).Update("complete", true)
	if res.Error != nil {
		return fmt.Errorf("failed to update task with ID %d: %w", id, res.Error) // More context
	}
	fmt.Println("Your task has been marked as complete.")
	return nil
}

// List retrieves and prints all tasks with nicely formatted output.
// It returns an error if the database is not connected, or if the query fails.
func List() error {
	if err := checkDB(); err != nil {
		return err
	}
	db := Database.DB

	var tasks []Model.Todo
	if res := db.Find(&tasks); res.Error != nil {
		return fmt.Errorf("failed to retrieve tasks: %w", res.Error)
	}

	maxIDWidth, maxTaskWidth := calculateColumnWidths(tasks)

	// Create and print the table header.
	printTable(tasks, maxIDWidth, maxTaskWidth)

	return nil
}

// calculateColumnWidths calculates the maximum widths required for the ID and Task columns.
func calculateColumnWidths(tasks []Model.Todo) (int, int) {
	maxIDWidth := len(idHeader)
	maxTaskWidth := len(taskHeader)

	for _, task := range tasks {
		idLength := len(strconv.Itoa(int(task.ID)))
		if idLength > maxIDWidth {
			maxIDWidth = idLength
		}
		taskLength := len(task.Task)
		if taskLength > maxTaskWidth {
			maxTaskWidth = taskLength
		}
	}
	return maxIDWidth, maxTaskWidth
}

// createSeparatorLine creates a horizontal line for the table.
func createSeparatorLine(idWidth, taskWidth int) string {
	totalWidth := idWidth + taskWidth + len(stateHeader) + (separatorWidth * 3) // 3 separators: | ID | Task | State |
	return strings.Repeat("-", totalWidth+3)
}

// createTaskLine formats a single task into a table row.
func createTaskLine(task Model.Todo, idWidth, taskWidth int) string {
	difresntStateSize := 0
	state := incompleteState
	if task.Complete {
		state = completeState
		difresntStateSize = 5
	}

	return fmt.Sprintf("| %-*d | %-*s | %-*s |", idWidth, task.ID, taskWidth, task.Task, len(stateHeader)+difresntStateSize, state)
}

// createHeader creates the header line of our table
func createHeader(idWidth, taskWidth int) string {
	return fmt.Sprintf("| %-*s | %-*s | %-*s |", idWidth, idHeader, taskWidth, taskHeader, len(stateHeader)+5, stateHeader)
}

// printTable takes the tasks and max widths, then prints them to standard output.
func printTable(tasks []Model.Todo, maxIDWidth, maxTaskWidth int) {

	separatorLine := createSeparatorLine(maxIDWidth, maxTaskWidth)

	fmt.Println(separatorLine)
	fmt.Println(createHeader(maxIDWidth, maxTaskWidth))
	fmt.Println(separatorLine)

	for _, task := range tasks {
		fmt.Println(createTaskLine(task, maxIDWidth, maxTaskWidth))
	}
	fmt.Println(separatorLine)
}

func Delete(arg string) error {
	id, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	//if err := checkDB(); err != nil {
	//	return err
	//}
	db := Database.DB

	res := db.Where("id = ?", id).Delete(&Model.Todo{})
	if res.Error != nil {
		return fmt.Errorf("failed to delete task with ID %d: %w", id, res.Error)
	}
	fmt.Println("Your task has been deleted.")
	return nil
}
