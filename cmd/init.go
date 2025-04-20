package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize taskman workspace",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting user home directory:", err)
			return
		}

		taskDir := filepath.Join(homeDir, ".taskman")
		todoFile := filepath.Join(taskDir, "todos.csv")

		err = os.MkdirAll(taskDir, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}

		if _, err := os.Stat(todoFile); os.IsNotExist(err) {
			_, err := os.Create(todoFile)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}

			fmt.Println("🎉 taskman initialized at", taskDir)
		} else {
			fmt.Println("✅ Already initialized:", todoFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
