package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as done",
	Long: `The done command takes an ID as input and marks the corresponding task as done.
This will update the status of the task in the todo list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the ID of the task to mark as done")
			return
		}

		idToMark := args[0]

		file, err := os.OpenFile("todos.csv", os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		rows, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Error reading CSV:", err)
			return
		}

		updated := false
		for i, row := range rows {
			if row[0] == idToMark {
				rows[i][1] = "done"
				updated = true
				break
			}
		}

		if !updated {
			fmt.Printf("Task with ID %v not found\n", idToMark)
			return
		}

		file.Truncate(0)
		file.Seek(0, 0)

		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := writer.WriteAll(rows); err != nil {
			fmt.Println("Error writing to CSV:", err)
			return
		}

		fmt.Printf("✅ Marked task %s as done\n", idToMark)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
