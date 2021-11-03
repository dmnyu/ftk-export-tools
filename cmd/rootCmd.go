package cmd

import "github.com/spf13/cobra"

var inputDirectory string

var rootCmd = &cobra.Command{
	Use: "cobra",
}

func Execute() error {
	return rootCmd.Execute()
}






