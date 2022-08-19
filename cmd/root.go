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

func Init(tr repository.TaskRepository) {
	rootCmd = NewRootCommand()
	// tsk ls
	rootCmd.AddCommand(NewLsCommand(tr))
	// tsk find
	rootCmd.AddCommand(NewFindCommand(tr))
	// tsk new
	rootCmd.AddCommand(NewNewCommand(tr))
	// tsk todo
	rootCmd.AddCommand(NewTodoCommand(tr))
	// tsk doing
	rootCmd.AddCommand(NewDoingCommand(tr))
	// tsk done
	rootCmd.AddCommand(NewDoneCommand(tr))
	// tsk mod
	rootCmd.AddCommand(NewModCommand(tr))
	// tsk rm
	rootCmd.AddCommand(NewRmCommand(tr))
	// tsk board
	rootCmd.AddCommand(NewBoardCommand(tr))
	// tsk cmt
	rootCmd.AddCommand(NewNoteCommand(tr))
	// tsk clean
	rootCmd.AddCommand(NewCleanCommand())
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
