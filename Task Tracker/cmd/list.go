package cmd

import (
    // "fmt"
    "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use:     "list",
    Aliases: []string{"list"},
    Short:   "list all of task",
    Long:    "list all of task",
    Args:    cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string)  {
        List()
    },
}

var listDone = &	cobra.Command{
    Use:     "list done",
    Aliases: []string{"list done"},
    Short:   "list all of task that are done",
    Long:    "list all of task that are done",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string)  {
        ListDone()
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
		rootCmd.AddCommand(listDone)
}