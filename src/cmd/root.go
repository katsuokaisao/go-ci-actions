package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(testCmd)
}

var rootCmd = &cobra.Command{
	Use: "",
}

func Execute() error {
	return rootCmd.Execute()
}

var testCmd = &cobra.Command{
	Use: "test",
	Run: func(cmd *cobra.Command, args []string) {
		println("Hello, World!")
	},
}
