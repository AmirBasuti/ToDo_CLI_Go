package TodoBehavior

import (
	"ToDo/internal/Database"
	"ToDo/internal/Model"
	"fmt"
)

func Add(taskTxt string) error {
	db := Database.DB
	task := Model.Todo{Task: taskTxt}
	if res := db.Create(&task); res.Error != nil {
		return fmt.Errorf("Create task failed\n %s \n\n", res.Error)
	}
	fmt.Printf("Your task was successfully created with this ID == %d \n\n", task.ID)
	return nil
}

func Complete(id uint) error {
	db := Database.DB
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	if res := db.Model(&Model.Todo{}).Where("id = ?", id).Update("complete", true); res.Error != nil {
		return fmt.Errorf("task with this ID %d failed to update \n\n %s", id, res.Error)
	}
	fmt.Println("Your task is complete")
	return nil
}
