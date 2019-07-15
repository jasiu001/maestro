package cli

import (
	"github.com/jasiu001/maestro/cli/cmd"
	"github.com/spf13/cobra"
	"log"
)

func RunMaestro(app cmd.Application, importer cmd.Importer) {
	var cliApplication = &cobra.Command{Use: "maestro"}
	cliApplication.AddCommand(cmd.ManualCommand)
	cliApplication.AddCommand(cmd.GenerateTranslateCommand(app))
	cliApplication.AddCommand(cmd.GenerateEnterCommand(importer))

	err := cliApplication.Execute()
	if err != nil {
		log.Fatalf("something went wrong during run cli version of application: %s", err)
	}
}
