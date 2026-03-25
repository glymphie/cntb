package cmd

import (
	"context"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var removeInstanceToFirewallCmd = &cobra.Command{
	Use:     "firewall [firewallId] [instanceId]",
	Short:   "remove instance to a firewall",
	Long:    `remove a specific instance to a specific firewall`,
	Example: `cntb unassign firewall 6bf66896-76db-4e32-9631-a77d3dc1c7f3 1001`,
	Run: func(cmd *cobra.Command, args []string) {
		_, httpResp, err := client.ApiClient().FirewallsApi.
			UnassignInstanceFirewall(context.Background(), unassignFirewallId, unassignInstanceId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while adding instance to firewall")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 2 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 2 {
			cmd.Help()
			log.Fatal("Please provide an firewallId and instanceId.")
		}

		unassignInstanceIdInt64, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		unassignFirewallId = args[0]
		unassignInstanceId = unassignInstanceIdInt64

		return nil
	},
}

func init() {
	contaboCmd.UnassignInstanceCmd.AddCommand(removeInstanceToFirewallCmd)
}
