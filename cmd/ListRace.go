package cmd

import (
	"github.com/spf13/cobra"
)

// ListRacesCmd represents the ListRace command
var ListRacesCmd = &cobra.Command{
	Use:     "List Races",
	Aliases: []string{"list"},
	Short:   "List all Races",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		ListRace()
	},
}

func init() {
	rootCmd.AddCommand(ListRacesCmd)
}
