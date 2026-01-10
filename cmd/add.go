package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command is used to add your tasks",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

tasktraker-cli add "First write function for add"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("%s\n", "please give me a description of the task")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
