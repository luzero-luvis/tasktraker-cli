package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"tesktracker-cli/internal"

	"github.com/spf13/cobra"
)

var markinprogessCmd = &cobra.Command{
	Use:   "markinprogess",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

./tasktraker-cli markinprogess <id> .`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("%s", "please provide the correct id")
			os.Exit(1)
		}

		id, err := strconv.Atoi((args[0]))
		if err != nil {
			fmt.Printf("%s", "please peovide the valid id")
		}

		tasks, err := internal.LoadTask()
		if err != nil {
			fmt.Printf("%s", "error loading task")
		}

		found := false
		for i, task := range tasks {
			if task.ID == id {
				found = true
				tasks[i].Status = "in-progress"
				tasks[i].UpdatedAt = time.Now()
				fmt.Printf("(ID :%v) marked in in-progress\n", id)
			}
		}

		if !found {
			fmt.Printf("%s", "the id your try to update is not found")
		}

		if err := internal.SaveTask(tasks); err != nil {
			fmt.Printf("%s", "errror saving the tasks")
		}
	},
}

func init() {
	rootCmd.AddCommand(markinprogessCmd)
}
