package cmd

import (
	"ToDo/internal/TodoBehavior"
	"fmt"
	"github.com/spf13/cobra"
)

var addcmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"ADD", "Add"},
	Short:   "you can add task hear ",
	Long:    "you can simply add task look like this:\n",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := TodoBehavior.Add(args[0]); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addcmd)
}
