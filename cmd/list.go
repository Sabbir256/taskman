package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

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

		if len(records) <= 1 {
			fmt.Println("No todos found.")
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Description", "Status"})
		table.SetBorder(false)
		table.SetHeaderLine(false)
		table.SetRowLine(false)
		table.SetColumnSeparator("")
		table.SetAutoWrapText(false)

		for i, row := range records {
			if i == 0 {
				continue
			}

			if statusFilter != "" && row[2] != statusFilter {
				continue
			}
			table.Append(row)
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVar(&statusFilter, "status", "", "Filter todos by status (Pending/Completed)")
}
