/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	taskCmd.AddCommand(addCmd)
}
