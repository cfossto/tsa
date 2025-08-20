/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCredsCmd represents the configCreds command
var configCredsCmd = &cobra.Command{
	Use:   "creds",
	Short: "Handles credentials for TSA (Tailscale Access Control CLI)",
	Long: `The creds command allows you to manage credentials for TSA (Tailscale Access Control CLI).
You can use this command to add, remove, or view credentials used for accessing Tailscale services.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configCreds called")
	},
}

func init() {
	configCmd.AddCommand(configCredsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCredsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCredsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
