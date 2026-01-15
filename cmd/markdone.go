package cmd

import (
	"fmt"
	"strconv"
	"time"

	"tesktracker-cli/internal"

	"github.com/spf13/cobra"
)

// markdoneCmd represents the markdone command
var markdoneCmd = &cobra.Command{
	Use:   "markdone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

./tasktracker-cli markdone <id>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("%s", "please use help command to check that usage of the command")
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("%s", "not a vaild id")
		}

		tasks, err := internal.LoadTask()
		if err != nil {
			fmt.Printf("%s", "eror loading task")
		}

		found := false
		for i, task := range tasks {
			if task.ID == id {
				found = true
				tasks[i].Status = "done"
				tasks[i].UpdatedAt = time.Now()
			}
		}

		if !found {
			fmt.Printf("%s", "that id you provided not found")
		}

		if err := internal.SaveTask(tasks); err != nil {
			fmt.Printf("%s", "error saving tasks")
		}
	},
}

func init() {
	rootCmd.AddCommand(markdoneCmd)
}
