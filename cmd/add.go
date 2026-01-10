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
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
