package cmd

import (
	"fmt"
	"github.com/dbraley/gophercises-task/db"
	"github.com/spf13/cobra"
	"strconv"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Failed to parse the argument %v:\n", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as complerted. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as complerted.\n", id)

			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
