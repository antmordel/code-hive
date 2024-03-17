package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "CLI Application",
}

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Print the current time",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current time: ", time.Now().Format("15:04:05"))
	},
}

func main() {
	rootCmd.AddCommand(timeCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

