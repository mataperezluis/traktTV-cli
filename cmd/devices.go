package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(devicesCmd)
}

var devicesCmd = &cobra.Command{
  Use:   "devices",
  Short: "initiates connection to traktTV API",
  Long:  `Device authentication is for apps and services with limited input or display capabilities`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("devices")
  },
}
