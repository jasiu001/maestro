package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ManualCommand = &cobra.Command{
	Use:   "manual",
	Short: "Shows all command with description/explenation",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		showDescription()
	},
}

func showDescription() {
	fmt.Println("TODO: in this place will be full description of commands application")
}
