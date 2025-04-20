package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type groupInfo struct {
	label string
	date time.Time
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `Displays all the tasks in the list with their ID and status.`,
	Run: func(cmd *cobra.Command, args []string) {
		const fileName = "todos.csv"

		file, err := os.Open(fileName)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("📋 No tasks found!")
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
			fmt.Println("📋 No tasks found!")
			return
		}

		groups := make(map[string][][]string)
		groupDates := make(map[string]time.Time)

		for _, row := range records {
			id, status, deadlineRaw, desc := row[0], row[1], row[2], row[3]

			groupLabel := humanizeDeadline(deadlineRaw)

			if _, ok := groups[groupLabel]; !ok {
				groups[groupLabel] = [][]string{}
				groupDates[groupLabel] = getParsedDeadline(deadlineRaw)
			}

			icon := getStatusIcon(status)
			groups[groupLabel] = append(groups[groupLabel], []string{icon, id, desc})
		}

		sortedGroups := getSortedGroups(groupDates)

		fmt.Println()
		for _, group := range sortedGroups {
			renderLabel(group.label)
			renderTasks(groups[group.label])
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func getStatusIcon(status string) string {
	if status == "done" {
		return "✔"
	}
	return " "
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
			return "Today"
		case 1:
			return "Tomorrow"
		default:
			if diff > 1 {
				return fmt.Sprintf("In %d days", int(diff))
			} else {
				return "Overdue"
			}
	}
}

func getParsedDeadline(deadline string) time.Time {
	if deadline == "" {
		return time.Time{}
	} else {
		parsed, err := time.Parse("2006-01-02", deadline)
		if err != nil {
			return time.Time{}
		} else {
			return parsed
		}
	}
}

func getSortedGroups(groupDates map[string]time.Time) []groupInfo {
	var sortedGroups []groupInfo
	for label, date := range groupDates {
		sortedGroups = append(sortedGroups, groupInfo{label, date})
	}

	sort.Slice(sortedGroups, func(i, j int) bool {
		return sortedGroups[i].date.Before(sortedGroups[j].date)
	})

	return sortedGroups
}

func renderLabel(label string) {
	switch label {
		case "Today":
			fmt.Println("  \033[32m" + label + "\033[0m") // green
		case "Tomorrow":
			fmt.Println("  \033[35m" + label + "\033[0m") // magenta
		case "Overdue":
			fmt.Println("  \033[31m" + label + "\033[0m") // red
		default:
			fmt.Println("  \033[34m" + label + "\033[0m") // blue
	}
}

func renderTasks(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", "", ""})
	table.SetAutoFormatHeaders(false)
	table.SetBorder(false)
	table.SetHeaderLine(false)
	table.SetRowLine(false)
	table.SetColumnSeparator("")
	table.SetAutoWrapText(false)
	table.SetColumnColor(tablewriter.Colors{tablewriter.FgHiGreenColor}, tablewriter.Colors{tablewriter.FgHiYellowColor}, tablewriter.Colors{})

	for _, row := range data {
		table.Append(row)
	}

	table.Render()
}
