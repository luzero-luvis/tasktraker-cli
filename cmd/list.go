package cmd

import (
	"fmt"
	"os"
	"strconv"

	"tesktracker-cli/internal"

	"github.com/fatih/color"
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

		white := color.New(color.FgWhite).SprintFunc()
		fmt.Println(white("TASK"))
		fmt.Println("──────────────────────────────────────────────────────")

		if len(args) > 0 {
			id, err := strconv.Atoi(args[0])
			if err == nil {
				for _, task := range tasks {
					if task.ID == id {
						fmt.Printf("ID(%d) %s %s (%s) (%s)",
							task.ID,
							task.Description,
							ColorStatus(task.Status),
							task.CreatedAt.Format("2006-01-15 15:04:05"),
							task.UpdatedAt.Format("2006-01-15 15:04:05"),
						)
					}
				}
				return
			}

			filterStatus := args[0]
			for _, task := range tasks {
				if filterStatus == task.Status {
					fmt.Printf("ID(%d) %s %s (%s) (%s)",
						task.ID,
						task.Description,
						ColorStatus(task.Status),
						task.CreatedAt.Format("2006-01-15 15:04:05"),
						task.UpdatedAt.Format("2006-01-15 15:04:05"),
					)
				}
			}
		}
		for _, task := range tasks {
			printTask(task)
		}
	},
}

func printTask(task internal.Task) {
	yellow := color.New(color.FgYellow).SprintfFunc()
	fmt.Printf("ID(%d) %s %s (%s) (%s)",
		task.ID,
		task.Description,
		ColorStatus(task.Status),
		task.CreatedAt.Format("2006-01-15 15:04:05"),
		task.UpdatedAt.Format("2006-01-15 15:04:05"))
	fmt.Println()
	fmt.Println(yellow("──────────────────────────────────────────────────────"))
}

func ColorStatus(status string) string {
	red := color.New(color.FgHiRed).SprintfFunc()
	yellow := color.New(color.FgHiYellow).SprintfFunc()
	blue := color.New(color.FgHiBlue).SprintfFunc()

	switch status {
	case "todo":
		return red("todo")
	case "in-progress":
		return yellow("in-progress")
	case "done":
		return blue("done")
	default:
		fmt.Printf("%s", "status not found")
	}
	return status
}

func init() {
	rootCmd.AddCommand(listCmd)
}
