package cmd

import (
	"fmt"
	"os"
	"time"

	"tesktracker-cli/internal"

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
			os.Exit(1)
		}

		tasks, err := internal.LoadTask()
		if err != nil {
			fmt.Printf("error loading task %v\n", err)
			os.Exit(1)
		}

		for _, tasks := range tasks {
			if tasks.Description == args[0] {
				fmt.Printf("task already exists at %d\n", tasks.ID)
				os.Exit(1)
			}
		}

		newTask := internal.Task{
			ID:          len(tasks) + 1,
			Description: args[0],
			Completed:   false,
			CompletedAt: time.Now(),
		}
		tasks = append(tasks, newTask)

		if err := internal.SaveTask(tasks); err != nil {
			fmt.Printf("error saving task %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("new task save successfully %v\n", newTask.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
