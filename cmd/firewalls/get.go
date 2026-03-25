package cmd

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var firewallGetCmd = &cobra.Command{
	Use:     "firewall [firewallId]",
	Short:   "Info about a specific firewall",
	Long:    `Retrieves information about one firewall by its id.`,
	Example: `cntb get firewall e5a94e71-fda6-4598-9e37-5b78d4e74b72`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().
			FirewallsApi.RetrieveFirewall(context.Background(), getFirewallId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving firewall")

		firewalls := make([]jmap, 0)
		for _, firewall := range resp.Data {
			updatedFirewall, _ := util.StructToMap(firewall)
			totalInbound := 0

			if firewall.Rules.Inbound != nil {
				totalInbound = len(firewall.Rules.Inbound)
			}

			updatedFirewall["noOfRules"] = totalInbound
			updatedFirewall["noOfInstances"] = len(firewall.Instances)

			var instances []string
			for _, instance := range firewall.Instances {
				instances = append(instances, strconv.FormatInt(instance.InstanceId, 10))
			}
			if instances == nil {
				updatedFirewall["instances"] = "[]"
			} else {
				updatedFirewall["instances"] = "[" + strings.Join(instances, ", ") + "]"
			}

			firewalls = append(firewalls, updatedFirewall)
		}

		responseJson, _ := json.Marshal(firewalls)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"firewallId", "name", "status", "noOfInstances", "noOfRules", "instances", "description",
			},
			WideFilter: []string{
				"firewallId", "name", "status", "noOfInstances", "noOfRules", "instances", "description", "createdDate",
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

		getFirewallId = args[0]

		if getFirewallId == "" {
			cmd.Help()
			log.Fatal("Argument firewallId is empty. Please provide a non empty firewallId.")
		}

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(firewallGetCmd)
}
