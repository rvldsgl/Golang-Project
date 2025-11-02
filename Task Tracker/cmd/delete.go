package cmd

import (
    // "fmt"
    "github.com/spf13/cobra"
)
var deleteCmd = &cobra.Command{
    Use:     "delete",
    Aliases: []string{"delete"},
    Short:   "delete task",
    Long:    "delete task",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string)  {
        delete(args[0])
    },
}


func init() {
    rootCmd.AddCommand(deleteCmd)
}