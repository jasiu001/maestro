package cmd

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
	"io"
	"log"
	"strings"
)

const exit = "exit!"

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

func exitProgram(input string, inputError error) bool {
	switch {
	case inputError == readline.ErrInterrupt:
		if len(input) == 0 {
			return true
		}
	case inputError == io.EOF:
	case strings.HasPrefix(input, exit):
		return true
	}

	return false
}

func readLine(shell *readline.Instance) (string, error) {
	line, err := shell.Readline()
	input := strings.TrimSpace(line)

	return input, err
}

func deferProgram(shell *readline.Instance) {
	err := shell.Close()
	if err != nil {
		log.Fatalf("failed during close read liner: %s", err)
	}
}
