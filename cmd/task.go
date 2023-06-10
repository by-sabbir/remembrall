/*
Copyright © 2023 NAME HERE <me@sabbir.dev>
*/
package cmd

import (
	"context"

	"github.com/by-sabbir/remembrall/db"
	"github.com/by-sabbir/remembrall/internal/task"
	v1 "github.com/by-sabbir/remembrall/internal/types/v1"
	"github.com/rivo/tview"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manages task",
	Long: `
		add - adds a new task
		remove - removes a task
		state - manages task state: todo, in-progress, and done
	`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := getTasks()
		if err != nil {
			log.Error("list not found: ", err)
		}
		app := tview.NewApplication()
		list := tview.NewList().ShowSecondaryText(false)

		for i, task := range tasks {

			list.AddItem(task.Title, "", rune(97+i), func() {
				idx := list.GetCurrentItem()
				text := tasks[idx].Title + "\t\t" + tasks[idx].Status
				list.SetItemText(list.GetCurrentItem(), text, "")
			})
		}
		list.SetBorder(true).SetTitle("All Tasks").SetTitleAlign(tview.AlignCenter)
		if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)
}

func getTasks() ([]v1.Task, error) {
	ctx := context.TODO()
	db, err := db.NewDBClient()
	if err != nil {
		log.Error("db initialization failed: ", err)
		return []v1.Task{}, err
	}

	svc := task.NewTaskService(db)
	tasks, err := svc.ListTask(ctx)
	if err != nil {
		log.Error("error returning list: ", err)
		return []v1.Task{}, err
	}
	return tasks, nil
}
