/*
Copyright © 2023 NAME HERE <me@sabbir.dev>
*/
package cmd

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/by-sabbir/remembrall/db"
	"github.com/by-sabbir/remembrall/internal/task"
	v1 "github.com/by-sabbir/remembrall/internal/types/v1"
	"github.com/gdamore/tcell/v2"
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
		table := tview.NewTable().SetBorders(true)

		headers := strings.Split("# Title Status Time-Remaining", " ")

		for i, header := range headers {
			table.SetCell(0, i,
				tview.NewTableCell(header).
					SetTextColor(tcell.ColorGreen).
					SetAlign(tview.AlignCenter).
					SetExpansion(100))
		}

		for i, task := range tasks {
			remainingTime := task.Deadline - time.Since(task.CreatedAt)
			id := strconv.Itoa(int(task.ID))
			rowList := []string{
				id, task.Title, task.Status, remainingTime.Truncate(1 * time.Minute).String(),
			}
			for c, item := range rowList {
				table.SetCell(i+1, c,
					tview.NewTableCell(item).
						SetTextColor(tcell.ColorWhite).
						SetAlign(tview.AlignCenter).
						SetExpansion(100))
			}
		}
		table.SetCell(len(tasks)+2, 3,
			tview.NewTableCell("Press ESC to exit.").
				SetTextColor(tcell.ColorYellow).
				SetAlign(tview.AlignCenter).
				SetExpansion(100))
		table.Select(1, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEscape {
				app.Stop()
			}
			if key == tcell.KeyEnter {
				table.SetSelectable(true, true)
			}
		}).SetSelectedFunc(func(row int, column int) {
			table.GetCell(row, 1).SetTextColor(tcell.ColorYellow)
			table.SetSelectable(true, true)
		})
		if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
			log.Error("could not build tui: ", err)
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
