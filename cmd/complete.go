package cmd

import (
	"ToDo/internal/TodoBehavior"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var completecmd = &cobra.Command{
	Use:     "Complete",
	Aliases: []string{"complete", "comp"},
	Short:   "Mark a task as complete.",
	Long:    "Mark a task as complete by providing its ID:\n",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if id, err := strconv.ParseUint(args[0], 10, 64); err == nil {
			fmt.Println(id)
			if err := TodoBehavior.Complete(uint(id)); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completecmd)
}
