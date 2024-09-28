/*
Copyright © 2022 github.com/Nathene
*/
package cmd

import (
	"os"

	"github.com/Nathene/toolbox/cmd/net"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "toolbox",
		Short: "toolbox app",
		Long:  "this is a toolbox applicaation, i dont know where ill be going with this.",
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	addSubCommand(rootCmd, net.NetCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "help message for toggle")
}

func addSubCommand(root *cobra.Command, follower *cobra.Command) {
	root.AddCommand(follower)
}
