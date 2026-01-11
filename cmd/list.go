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

tasktracker-cli list`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := internal.LoadTask()
		if err != nil {
			fmt.Printf("%s", "error listing task")
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("list is empty")
			os.Exit(1)
		}
		if len(args) > 0 {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("%s", "please provide valide id or id not exists")
			}

			for _, tasks := range tasks {
				if tasks.ID == id {
					fmt.Printf("%d\t", tasks.ID)
					fmt.Printf("%s\t", tasks.Description)
					fmt.Printf("%t\t", tasks.Completed)
					fmt.Printf("%v\n", tasks.CompletedAt)
					return
				}
			}
		}
		for _, tasks := range tasks {
			fmt.Printf("%d\t", tasks.ID)
			fmt.Printf("%s\t", tasks.Description)
			fmt.Printf("%t\t", tasks.Completed)
			fmt.Printf("%v\n", tasks.CompletedAt)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
