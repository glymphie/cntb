package cmd

import (
	"context"
	"encoding/json"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// historyCmd represents the history command
var firewallHistoryCmd = &cobra.Command{
	Use:     "firewalls",
	Short:   "History of your firewalls",
	Long:    `Show what happend to your firewalls over time.`,
	Example: `cntb history firewalls`,
	Run: func(cmd *cobra.Command, args []string) {
		historyRequest := client.ApiClient().FirewallsAuditsApi.
			RetrieveFirewallAuditsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if historyfirewallIdFilter != "" {
			historyRequest = historyRequest.FirewallId(historyfirewallIdFilter)
		}

		resp, httpResp, err := historyRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving firewalls history")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"id", "firewallId", "action", "username", "timestamp",
			},
			WideFilter: []string{
				"id", "firewallId", "action", "username", "changedBy",
				"requestId", "traceId", "timestamp", "changes",
			},
			JsonPath: contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("firewallId", cmd.Flags().Lookup("firewallId"))
		historyfirewallIdFilter = viper.GetString("firewallId")

		return nil
	},
}

func init() {
	contaboCmd.HistoryCmd.AddCommand(firewallHistoryCmd)

	firewallHistoryCmd.Flags().StringVarP(&historyfirewallIdFilter, "firewallId", "", "",
		`Filter by a specific firewall via its firewallId.`)
}
