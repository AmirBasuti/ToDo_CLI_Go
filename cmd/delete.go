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
	Short:   "you can add task hear ",
	Long:    "you can simply add task look like this:\n",
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
