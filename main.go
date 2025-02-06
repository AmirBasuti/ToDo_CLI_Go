/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"ToDo/cmd"
	"ToDo/internal/Database"
	"fmt"
)

func main() {
	//Database.Connection()
	err := Database.Connection()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Execute()
}
