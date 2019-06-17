package cmd

import (
	"fmt"
	"log"

	"github.com/chzyer/readline"
	"github.com/jasiu001/maestro/cli/cmd/enter"
	"github.com/spf13/cobra"
)

type Importer interface {
	SetType(string)
	SetWords([]string)
	SetTranslation(string)
	SetDescription(string)
	Save() error
}

func GenerateEnterCommand(importer Importer) *cobra.Command {
	return &cobra.Command{
		Use:   "enter",
		Short: "Enter new words to translate",
		Long:  fmt.Sprintf("To exit press %q", exit),
		Run: func(cmd *cobra.Command, args []string) {
			runEnter(importer)
		},
	}
}

func runEnter(importer Importer) {
	shell, err := readline.NewEx(&readline.Config{
		Prompt:          "> ",
		InterruptPrompt: "^C",
		EOFPrompt:       exit,
	})
	if err != nil {
		log.Fatalf("failed during create new read liner: %s", err)
		return
	}
	defer deferProgram(shell)

	for {
		inputData := enter.NewInputData()

		// read words
		fmt.Println("Input the words. To end up press '/'")
		counter := 0
		keepGoing := false
		for {
			if keepGoing {
				break
			}
			input, err := readLine(shell)
			if exitProgram(input, err) {
				fmt.Println("Good bye!")
				return
			}
			if input == "/" {
				if counter == 0 {
					fmt.Println("At least one word is required")
					continue
				}
				keepGoing = true
				continue
			}
			err = inputData.SetWord(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
			counter++
		}

		// read translation/description of the words entered above
		for {
			element, is := inputData.GetElement()
			if !is {
				break
			}

			keepGoing = false
			for {
				if keepGoing {
					break
				}
				fmt.Printf("%s > \n", element.Label)

				input, err := readLine(shell)
				if exitProgram(input, err) {
					fmt.Println("Good bye!")
					return
				}
				err = element.SetValue(input)
				if err != nil {
					fmt.Println(err)
					continue
				}
				keepGoing = true
			}
		}

		// repeat or exit
		keepGoing = false
		for {
			if keepGoing {
				break
			}
			fmt.Println("Next? [Y/n]: ")

			input, err := readLine(shell)
			if exitProgram(input, err) {
				fmt.Println("Good bye!")
				return
			}

			switch input {
			case "Y":
				saveInput(importer, inputData)
				keepGoing = true
				break
			case "n":
				saveInput(importer, inputData)
				return
			default:
				fmt.Println("> incorrect expression. Try again.")
			}
		}
	}
}

type InputReceiver interface {
	GetType() string
	GetWords() []string
	GetTranslation() string
	GetDescription() string
}

func saveInput(importer Importer, data InputReceiver) {
	importer.SetType(data.GetType())
	importer.SetWords(data.GetWords())
	importer.SetTranslation(data.GetTranslation())
	importer.SetDescription(data.GetDescription())

	err := importer.Save()
	if err != nil {
		fmt.Printf("failed during save words: %s. Please try again.", err)
		return
	}

	fmt.Println("Words are saved")
}
