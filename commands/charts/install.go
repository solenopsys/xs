package charts

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"solenopsys-cli-xs/utils"
)

var cmdInstall = &cobra.Command{
	Use:   "install [chart] [version] [repository]",
	Short: "Install chart",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utils.GetConfig()
		if err != nil {
			log.Fatal(err)
		}
		api := utils.NewAPI(config)
		chart := args[0]
		repoUrl := args[2]
		version := args[1]
		api.CreateHelmChartSimple(chart, repoUrl, version)

		fmt.Println("Installed: ", chart)
	},
}