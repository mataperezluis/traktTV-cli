package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of traktTV-cli",
  Long:  `All software has versions. This is traktTV-cli's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("traktTV-cli v0.1")
  },
}
