package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
    Use:     "update",
    Aliases: []string{"update"},
    Short:   "update your task in todo list",
    Long:    "update your task in todo lis",
    Args:    cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string)  {
        fmt.Printf("%s.\n", Update(args[0], args[1]))
    },
}

var inProgressCmd = &cobra.Command{
    Use:     "mark-in-progress",
    Aliases: []string{"mark-in-progress"},
    Short:   "update to mark-in-progress",
    Long:    "update to mark-in-progress",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string)  {
        fmt.Printf("%s.\n", MarkInProgress(args[0]))
    },
}

var doneCmd = &cobra.Command{
    Use:     "mark-done",
    Aliases: []string{"mark-done"},
    Short:   "mark-done",
    Long:    "mark-done",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string)  {
        fmt.Printf("%s.\n", MarkDone(args[0]))
    },
}

func init() {
    rootCmd.AddCommand(updateCmd)
		rootCmd.AddCommand(inProgressCmd)
		rootCmd.AddCommand(doneCmd)
}