package cmd

import (
	"log"

	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "tsk",
		Short: "tsk is a cli task management tool",
	}
}

func Init(tr *repository.TaskRepository) {
	rootCmd = NewRootCommand()
	rootCmd.AddCommand(NewLsCommand(tr))
	rootCmd.AddCommand(NewFindCommand(tr))
	rootCmd.AddCommand(NewNewCommand(tr))
	rootCmd.AddCommand(NewModCommand(tr))
	rootCmd.AddCommand(NewRmCommand(tr))
	rootCmd.AddCommand(NewBoardCommand(tr))
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
