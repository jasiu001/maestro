package cmd

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const exit = "exit!"

type Application interface {
	GetDescription() string
	NumberOfWords() int
	ExecuteResponse([]string)
	GetResult() string
	IsFinished() bool
	GetSummary() string
}

func GenerateTranslateCommand(app Application) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Run translation exercise",
		Long:  fmt.Sprintf("To exit press %q", exit),
		Run: func(cmd *cobra.Command, args []string) {
			runTranslate(app)
		},
	}
}

func runTranslate(app Application) {
	shell, err := readline.NewEx(&readline.Config{
		Prompt: "> ",
	})
	if err != nil {
		log.Fatalf("failed during create new read liner: %s", err)
		return
	}
	defer func() {
		err := shell.Close()
		if err != nil {
			log.Fatalf("failed during close read liner: %s", err)
		}
	}()

	for {
		if app.IsFinished() {
			fmt.Printf("Exercise is finished. Summary: %s \n", app.GetSummary())
			fmt.Println("Exit")
			break
		}
		fmt.Println(app.GetDescription())
		var input []string
		for i := 0; i < app.NumberOfWords(); i++ {
			inp, _ := shell.Readline()
			if strings.HasPrefix(inp, exit) {
				fmt.Println("Good bye!")
				return
			}
			input = append(input, inp)
		}
		app.ExecuteResponse(input)
		fmt.Printf("Result: %s \n", app.GetResult())
	}
}