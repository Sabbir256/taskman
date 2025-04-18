package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var loud bool

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greets the user interactively",
	Long: `The greet command asks for your name and says hello`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your name: ")
		name, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		name = strings.TrimSpace(name)

		if loud {
			fmt.Printf("HELLO, %s!!!\n", strings.ToUpper(name))
		} else {
			fmt.Printf("Hello, %v!\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)

	greetCmd.Flags().BoolVarP(&loud, "loud", "l", false, "Print greeting in uppercase")
}
