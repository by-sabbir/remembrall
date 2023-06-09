/*
Copyright © 2023 NAME HERE <me@sabbir.dev>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/by-sabbir/remembrall/db"
	"github.com/by-sabbir/remembrall/internal/task"
	v1 "github.com/by-sabbir/remembrall/internal/types/v1"
	"github.com/google/uuid"
	"github.com/rivo/tview"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",

	Run: func(cmd *cobra.Command, args []string) {
		var t v1.Task
		t.ID = uuid.NewString()
		t.CreatedAt = time.Now()
		status := []string{"ToDo", "InProgress", "Done"}
		// t.Deadline = 15 * time.Minute
		// t.Title = strings.Join(args, " ")
		// t.Status = "InProgress"

		tui := tview.NewApplication()

		form := tview.NewForm().
			AddInputField("Title", "", 32, nil, func(title string) {
				t.Title = title
			}).AddDropDown("Status", status, 0, func(option string, _ int) {
			t.Status = option
		}).AddButton("Done!", func() {
			fmt.Println("Task Added: ", t)
			time.Sleep(1 * time.Second)
			tui.Stop()
		})
		err := tui.SetRoot(form, true).EnableMouse(false).Run()
		if err != nil {
			log.Error("could not run tui")
			os.Exit(1)
		}
		createdTask, err := taskAdd(&t)
		if err != nil {
			log.Error("failed to create task: ", err)
		} else {
			log.Info("got task: ", t)
			log.Info("task created successfully: ", createdTask)
		}

	},
}

func init() {
	taskCmd.AddCommand(addCmd)
}

func taskAdd(t *v1.Task) (*v1.Task, error) {

	ctx := context.TODO()
	db, err := db.NewDBClient()
	if err != nil {
		log.Error("db initialization failed: ", err)
		return &v1.Task{}, err
	}

	svc := task.NewTaskService(db)

	createdTask, err := svc.AddTask(ctx, t)
	if err != nil {
		log.Error("could not create task: ", err)
		return &v1.Task{}, err
	}

	return createdTask, nil
}
