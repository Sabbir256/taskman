package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var statusFilter string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long: `Displays all the todos in the list with their ID and status.`,
	Run: func(cmd *cobra.Command, args []string) {
		const fileName = "todos.csv"

		file, err := os.Open(fileName)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("No todos found.")
				return
			}
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		if len(records) == 0 {
			fmt.Println("No todos found.")
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		// table.SetHeader([]string{"Id", "Status"})
		table.SetBorder(false)
		table.SetHeaderLine(false)
		table.SetRowLine(false)
		table.SetColumnSeparator("")
		table.SetAutoWrapText(false)

		for _, row := range records {
			id, status, deadline, desc := row[0], row[1], row[2], row[3]

			if status == "pending" || status == "" {
				status = "[ ]"
			} else if status == "done" {
				status = "[✔]"
			}

			humanDeadline := humanizeDeadline(deadline)
			table.Append([]string{id, status, humanDeadline, desc})
		}

		fmt.Println()
		table.Render()
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVar(&statusFilter, "status", "", "Filter todos by status (Pending/Completed)")
}

func humanizeDeadline(deadline string) string {
	if deadline == "" {
		return "No deadline"
	}

	parsed, err := time.Parse("2006-01-02", deadline)
	if err != nil {
		return deadline
	}

	today := time.Now().Truncate(24 * time.Hour)
	diff := parsed.Sub(today).Hours() / 24

	switch diff {
		case 0:
			return "today"
		case 1:
			return "tomorrow"
		case -1:
			return "yesterday"
		default:
			return parsed.Format("Mon Jan 2")
	}
}
