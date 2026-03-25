package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var AssignInstanceCmd = &cobra.Command{
	Use:   "assign",
	Short: "Add instance to private network or firewall",
	Long:  `Add a specific instance to a specific private network using its ip or to a firewall using firewallId`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args:       cobra.OnlyValidArgs,
	SuggestFor: []string{"assign", "add", "addInstance"},
	ValidArgs:  []string{"privateNetwork", "firewall"},
}

func init() {
	rootCmd.AddCommand(AssignInstanceCmd)
}
