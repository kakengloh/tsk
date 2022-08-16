package cmd

import (
	"log"

	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tsk",
		Short: "tsk is a cli task management tool",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
}

func Init(tr *repository.TaskRepository) {
	rootCmd = NewRootCmd()
	rootCmd.AddCommand(NewLsCmd(tr))
	rootCmd.AddCommand(NewMkCmd(tr))
	rootCmd.AddCommand(NewEdCmd(tr))
	rootCmd.AddCommand(NewRmCmd(tr))
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
