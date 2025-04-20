package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/Sabbir256/taskman/utils"
	"github.com/spf13/cobra"
)

var deadline string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long: `The add command takes a description of the task and adds it to the list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a description for the task.")
			return
		}
		description := args[0]

		fileName := utils.GetTodoFilePath()
		var nextID int = 1

		// check if the file exists
		_, err := os.Stat(fileName)
		newFile := os.IsNotExist(err)

		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		records, _ := reader.ReadAll()

		if !newFile && len(records) > 0 {
			lastRecord := records[len(records)-1]
			lastID, err := strconv.Atoi(lastRecord[0])
			if err == nil {
				nextID = lastID + 1
			}
		}

		writer := csv.NewWriter(file)
		defer writer.Flush()

		record := []string{
			strconv.Itoa(nextID),
			"pending",
			deadline,
			description,
		}

		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record to file:", err)
			return
		}

		fmt.Println("✅ Successfully added a new task!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&deadline, "deadline", "", "Deadline for the task (e.g., 2025-04-11)")
}
