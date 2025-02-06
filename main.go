/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"ToDo/internal/Database"
	"fmt"
)

func main() {
	//cmd.Execute()
	err := Database.Connection()
	if err != nil {
		fmt.Println(err)
	}
}
