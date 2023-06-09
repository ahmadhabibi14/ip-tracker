package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "IPtracker CLI App",
		Long:  `IPtracker CLI App`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
