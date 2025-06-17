package cmd

import (
	"github.com/spf13/cobra"
)

// Global Variable
var (
	rootCmd = &cobra.Command{
		Use:   "IP-Tracker-CLI",
		Short: "IP Tracker Mini Application",
		Long:  `IP Tracker Mini Application Using Cobra Cli Library.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
