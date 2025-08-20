/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set or view configuration for TSA",
	Long: `The config command allows you to set or view the configuration for TSA (Tailscale Access Control CLI).
You can use this command to manage settings such as API keys, server addresses, and other preferences.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

var addCredentialCmd = &cobra.Command{
	Use:   "add-credential",
	Short: "Add a new credential to the TSA configuration",
	Long: `The add-credential command allows you to add a new credential to the TSA configuration.
When you invoke this command, you will be prompted to enter the details of the credential you wish to add.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add-credential called")
		// Here you would implement the logic to add a credential.
		// This could involve prompting the user for input or reading from a file.
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
