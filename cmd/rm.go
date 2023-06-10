/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"strconv"

	"github.com/by-sabbir/remembrall/db"
	"github.com/by-sabbir/remembrall/internal/task"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a task by id",

	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		db, err := db.NewDBClient()
		if err != nil {
			log.Error("db initialization failed: ", err)
		}

		svc := task.NewTaskService(db)

		for _, idStr := range args {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Error("could not parse int: ", err)
				continue
			}
			if err := svc.RemoveTask(ctx, id); err != nil {
				log.Error("failed to remove item: ", id, err)
			}
		}
	},
}

func init() {
	taskCmd.AddCommand(rmCmd)

}
