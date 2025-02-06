package cmd

import (
	"ToDo/internal/TodoBehavior"
	"fmt"
	"github.com/spf13/cobra"
)

var listcmd = &cobra.Command{
	Use:     "List",
	Aliases: []string{"list"},
	Short:   "you can complete task hear ",
	Long:    "you can simply add task look like this:\n",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := TodoBehavior.List()
		if err != nil {
			fmt.Println(err)

		}

	},
}

func init() {
	rootCmd.AddCommand(listcmd)
}
