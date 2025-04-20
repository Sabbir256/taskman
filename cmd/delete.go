package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/Sabbir256/taskman/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Command to delete a task",
	Long: `The delete command take an ID as input and deletes the corresponding task from the list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the ID of the task to delete.")
			return
		}

		idToDelete := args[0]

		fileName := utils.GetTodoFilePath()
		file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
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

		var updatedRows [][]string

		deleted := false
		for _, row := range rows {
			if row[0] == idToDelete {
				deleted = true
				continue
			}
			updatedRows = append(updatedRows, row)
		}

		if !deleted {
			fmt.Println("⚠️ Task ID not found!")
			return
		}

		file.Truncate(0)
		file.Seek(0, 0)

		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := writer.WriteAll(updatedRows); err != nil {
			fmt.Println("Error writing data to CSV:", err)
			return
		}

		fmt.Printf("🗑️ Deleted task %s from the list.\n", idToDelete)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
