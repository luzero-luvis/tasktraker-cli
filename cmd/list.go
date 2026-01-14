package cmd

import (
	"fmt"
	"os"
	"strconv"

	"tesktracker-cli/internal"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

./tasktracker-cli list
./tasktracker-cli list todo
./tasktracker-cli list in-progess
./tasktracker-cli list completed`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := internal.LoadTask()
		if err != nil {
			fmt.Printf("%s", "error listing task")
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("no task to do")
			os.Exit(1)
		}
		fmt.Println("ID | Description | Status       | CreatedAt")
		fmt.Println("---+-------------+--------------+---------------------------")
		if len(args) > 0 {
			arg := args[0]
			id, err := strconv.Atoi(args[0])
			if err != nil {
				if arg != "todo" && arg != "done" && arg != "in-progess" {
					fmt.Printf("%s", "please provide valide id or id not exists")
				}
			}

			for _, tasks := range tasks {
				if tasks.ID == id {
					fmt.Printf("%d  | %s  | %v  | %v  | %v\n",
						tasks.ID,
						tasks.Description,
						tasks.Status,
						tasks.CreatedAt,
						tasks.UpdatedAt)
					return
				}
			}
		} else {
			for _, tasks := range tasks {
				fmt.Printf("%d  | %s  | %v  | %v  | %v\n",
					tasks.ID,
					tasks.Description,
					tasks.Status,
					tasks.CreatedAt,
					tasks.UpdatedAt)
			}
		}

		if len(args) == 0 {
			return
		}

		filterSatus := args[0]
		for _, tasks := range tasks {
			if tasks.Status == filterSatus {
				fmt.Printf("%d | %s | %v | %v | %v |\n", tasks.ID, tasks.Description, tasks.Status, tasks.CreatedAt, tasks.UpdatedAt)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
