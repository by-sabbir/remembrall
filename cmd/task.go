/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/rivo/tview"
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
		fmt.Println("task called")
		app := tview.NewApplication()
		form := tview.NewForm().
			AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
			AddInputField("First name", "", 20, nil, nil).
			AddInputField("Last name", "", 20, nil, nil).
			AddTextArea("Address", "", 40, 0, 0, nil).
			AddTextView("Notes", "This is just a demo.\nYou can enter whatever you wish.", 40, 2, true, false).
			AddCheckbox("Age 18+", false, nil).
			AddPasswordField("Password", "", 10, '*', nil).
			AddButton("Save", nil).
			AddButton("Quit", func() {
				app.Stop()
			})
		form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
		if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)

}
