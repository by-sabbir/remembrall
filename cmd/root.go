/*
Copyright © 2023 NAME HERE <me@sabbir.dev>
*/
package cmd

import (
	"os"

	"github.com/by-sabbir/remembrall/db"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "remembrall",
	Short: "CLI based task manager with time-based alert",
	Long:  `Remembrall is the magical glass ball that turns red when you forget something!`,

	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	// initiate and migrate database
	d, err := db.NewDBClient()
	if err != nil {
		log.Error("could not initiate database: ", err)
		os.Exit(120)
	}
	if err := d.Migrate(); err != nil {
		log.Warn("could not migrate database: ", err)
	}
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
