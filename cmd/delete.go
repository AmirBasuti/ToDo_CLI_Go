package cmd

import (
	"ToDo/internal/TodoBehavior"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var deletecmd = &cobra.Command{
	Use:     "Delete",
	Aliases: []string{"delete", "del"},
	Short:   "Delete a task by its ID",
	Long:    "Use this command to delete a task by providing its ID. For example:\n",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if id, err := strconv.ParseUint(args[0], 10, 64); err == nil {
			if err := TodoBehavior.Delete(uint(id)); err != nil {
				fmt.Println(err)
			}
			TodoBehavior.List()
		} else {
			fmt.Println(err)

		}
	},
}

func init() {
	rootCmd.AddCommand(deletecmd)
}
