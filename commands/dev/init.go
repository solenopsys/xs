package dev

import (
	"github.com/spf13/cobra"
	"xs/services"
)

var cmdInit = &cobra.Command{
	Use:   "init ",
	Short: "Init tree repository",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		loader := services.NewLoader()
		loader.LoadBase()
		services.NewHelper().InitRepository()
		loader.SyncModules()
	},
}
