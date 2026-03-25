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
)

var firewallRulesGetCmd = &cobra.Command{
	Use:     "firewall-rules [firewallId]",
	Short:   "Info about rules of a firewall",
	Long:    `Retrieves information about rules of one firewall by its id.`,
	Example: `cntb get firewall-rules e5a94e71-fda6-4598-9e37-5b78d4e74b72`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().
			FirewallsApi.RetrieveFirewall(context.Background(), getFirewallRulesId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving firewall rules")

		rules := make([]jmap, 0)

		for _, firewall := range resp.Data {
			for _, ruleInbound := range firewall.Rules.Inbound {
				rule := make(jmap)

				rule["action"] = ruleInbound.Action
				rule["protocol"] = ruleInbound.Protocol
				rule["status"] = ruleInbound.Status

				if ruleInbound.DestPorts == nil {
					rule["destPorts"] = []string{}
				} else {
					rule["destPorts"] = ruleInbound.DestPorts
				}

				if ruleInbound.SrcCidr.Ipv4 == nil {
					rule["SrcCidrIpv4"] = []string{}
				} else {
					rule["SrcCidrIpv4"] = ruleInbound.SrcCidr.Ipv4
				}

				if ruleInbound.SrcCidr.Ipv6 == nil {
					rule["SrcCidrIpv6"] = []string{}
				} else {
					rule["SrcCidrIpv6"] = ruleInbound.SrcCidr.Ipv6
				}

				rules = append(rules, rule)
			}
		}

		responseJson, _ := json.Marshal(rules)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"action", "protocol", "destPorts", "status", "SrcCidrIpv4", "SrcCidrIpv6",
			},
			WideFilter: []string{
				"action", "protocol", "destPorts", "status", "SrcCidrIpv4", "SrcCidrIpv6",
			},
			JsonPath: contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an firewallId.")
		}

		getFirewallRulesId = args[0]

		if getFirewallRulesId == "" {
			cmd.Help()
			log.Fatal("Argument firewallId is empty. Please provide a non empty firewallId.")
		}

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(firewallRulesGetCmd)
}
