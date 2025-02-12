package cmd

import (
	"github.com/spf13/cobra"
)

// scheduleCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Lists the full season race schedule",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		schedule()
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)

}
