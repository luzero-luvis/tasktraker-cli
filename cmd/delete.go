package cmd

import (
	"fmt"
	"os"
	"strconv"

	"tesktracker-cli/internal"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task by ID",
	Long:  `Delete a task from the list. Usage: ./tasktraker-cli delete <id>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("%s", "please provide the valid id")
			os.Exit(1)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("%s", "invalid task id")
		}

		tasks, err := internal.LoadTask()
		if err != nil {
			fmt.Printf("%s", "error loading tasks")
		}

		found := false
		var UpdatedTask []internal.Task
		for _, tasks := range tasks {
			if tasks.ID == id {
				found = true
			} else {
				UpdatedTask = append(UpdatedTask, tasks)
			}
		}

		if !found {
			fmt.Printf("%s", "that id you are searching is not found")
		}

		if err := internal.SaveTask(UpdatedTask); err != nil {
			fmt.Printf("error saving tasks %v\n", err)
		} else {
			fmt.Printf("%s\n", "task deleted successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
