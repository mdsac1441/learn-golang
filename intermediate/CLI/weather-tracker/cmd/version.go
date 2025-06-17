package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the version",
	Long:  `Get latest the version using Command line utility.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(" Weather-CLI: v1.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
