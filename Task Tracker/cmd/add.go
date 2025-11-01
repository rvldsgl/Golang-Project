package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use:     "add",
    Aliases: []string{"addition"},
    Short:   "Add task into your todo list",
    Long:    "Add task into your todo list",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string)  {
        fmt.Printf("%s.\n", Add(args[0]))
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}