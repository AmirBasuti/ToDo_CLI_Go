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
	Short:   "you can complete task hear ",
	Long:    "you can simply add task look like this:\n",
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
