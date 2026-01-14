package cmd

import (
	"fmt"
	"strconv"
	"time"

	"tesktracker-cli/internal"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Printf("%s", "please give me an id and new description")
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("%s", "please provide correct id")
		}

		tasks, err := internal.LoadTask()
		if err != nil {
			fmt.Printf("%s", "error loading task")
		}

		found := false
		for i, task := range tasks {
			if task.ID == id {
				found = true
				_ = found
				tasks[i].Description = args[1]
				tasks[i].UpdatedAt = time.Now()
				break
			}
		}

		if !found {
			fmt.Printf("%s\n", "id not found")
		}

		if err := internal.SaveTask(tasks); err != nil {
			fmt.Printf("%s\n", "error saving the tasks")
		}

		fmt.Printf("%s\n", "task update successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
