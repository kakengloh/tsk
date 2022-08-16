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
	}
}

func Init(tr *repository.TaskRepository) {
	rootCmd = NewRootCmd()

	lsCmd := NewLsCmd(tr)
	rootCmd.AddCommand(lsCmd)

	mkCmd := NewMkCmd(tr)
	rootCmd.AddCommand(mkCmd)

	edCmd := NewEdCmd(tr)
	rootCmd.AddCommand(edCmd)

	rmCmd := NewRmCmd(tr)
	rootCmd.AddCommand(rmCmd)

	// Root command is an alias of `ls`
	rootCmd.RunE = lsCmd.RunE
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
