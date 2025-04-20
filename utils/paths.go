package utils

import (
	"os"
	"path/filepath"
)

func GetTodoFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(homeDir, ".taskman", "todos.csv")
}
