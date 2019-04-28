package cli

import (
	"github.com/jasiu001/maestro/cli/cmd"
	"github.com/spf13/cobra"
	"log"
)

func RunCmdApplication(app cmd.Application) {
	var cliApplication = &cobra.Command{Use: "app"}
	cliApplication.AddCommand(cmd.ManualCommand)
	cliApplication.AddCommand(cmd.GenerateTranslateCommand(app))

	err := cliApplication.Execute()
	if err != nil {
		log.Fatalf("something went wrong during run cli version of application: %s", err)
	}
}
