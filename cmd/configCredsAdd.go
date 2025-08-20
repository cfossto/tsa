/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCredsAddCmd represents the configCredsAdd command
var configCredsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add API token to your local machine",
	Long: `The add command allows you to add an API token to your local machine.
Will default to ~/.tsacl/ folder if no folder is specified.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configCredsAdd called")
	},
}

func init() {
	configCredsCmd.AddCommand(configCredsAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCredsAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCredsAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
