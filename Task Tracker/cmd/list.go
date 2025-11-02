package cmd

import (
    // "fmt"
    "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Long:  "List all of the tasks",
    Run: func(cmd *cobra.Command, args []string) {
        List()
    },
}

var listDoneCmd = &cobra.Command{
    Use:   "done",
    Short: "List all completed tasks",
    Long:  "List all tasks that are marked as done",
    Run: func(cmd *cobra.Command, args []string) {
        ListDone()
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
    listCmd.AddCommand(listDoneCmd) 
}