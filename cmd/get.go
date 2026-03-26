package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Display one or more resources",
	Long:  `Display one or more resources from the Kubernetes cluster. Use subcommands to specify the resource type.`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
