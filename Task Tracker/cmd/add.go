package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use:     "add",
    Aliases: []string{"addition"},
    Short:   "Add 2 numbers",
    Long:    "Carry out addition operation on 2 numbers",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string)  {
        fmt.Printf("%s.\n", Add(args[0]))
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}