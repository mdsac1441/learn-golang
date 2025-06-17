package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Check for test",
	Long:  "Test command line utility using Cobra library",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(" good test hoga bhai ")
		} else {
			fmt.Println("Test tha na bhai")
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
