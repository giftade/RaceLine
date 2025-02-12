package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// standingsCmd represents the standings command
var standingsCmd = &cobra.Command{
	Use:   "standings",
	Short: "Show the latest Driver and Constructors standings",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("standings called")
	},
}

func init() {
	rootCmd.AddCommand(standingsCmd)

	
}
