package cmd

import (
	"context"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var firewallDeleteCmd = &cobra.Command{
	Use:   "firewall [firewallId]",
	Short: "Deletes a specific firewall by id",
	Long:  `Specify a firewall id to delete. All the instances must be unassigned before deleting the firewall`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().
			FirewallsApi.DeleteFirewall(context.Background(), deleteFirewallId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting firewall")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide a firewallId")
		}

		deleteFirewallId = args[0]

		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(firewallDeleteCmd)
}
