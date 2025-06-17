package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ahmed",
		Short: "Weather application",
		Long:  "A simple weather Application using cobra library",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
