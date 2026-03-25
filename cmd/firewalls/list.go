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
	"github.com/spf13/viper"
)

type jmap map[string]interface{}

var firewallsGetCmd = &cobra.Command{
	Use:     "firewalls",
	Short:   "All about your firewalls",
	Long:    `Retrieves information about one or multiple firewalls. Filter by name.`,
	Example: `cntb get firewalls`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveFirewallListRequest := client.ApiClient().
			FirewallsApi.RetrieveFirewallList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if listFirewallNameFilter != "" {
			ApiRetrieveFirewallListRequest = ApiRetrieveFirewallListRequest.Name(listFirewallNameFilter)
		}

		if strings.TrimSpace(listFirewallInstancesFilter) != "" {
			parts := strings.Split(listFirewallInstancesFilter, ",")
			ids := make([]string, 0, len(parts))
			for _, p := range parts {
				s := strings.TrimSpace(p)
				if s != "" {
					ids = append(ids, s)
				}
			}
			if len(ids) > 0 {
				ApiRetrieveFirewallListRequest = ApiRetrieveFirewallListRequest.InstanceIds(strings.Join(ids, ","))
			}
		}

		resp, httpResp, err := ApiRetrieveFirewallListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving firewalls")

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

			if listFirewallIdFilter != "" &&
				firewall.FirewallId != listFirewallIdFilter {
				continue
			}

			if listFirewallStatusFilter != "" &&
				string(firewall.Status) != listFirewallStatusFilter {
				continue
			}

			firewalls = append(firewalls, updatedFirewall)
		}

		responseJson, _ := json.Marshal(firewalls)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"firewallId", "name", "status", "noOfInstances", "noOfRules", "instances", "description"},
			WideFilter: []string{"firewallId", "name", "status", "noOfInstances", "noOfRules", "instances", "description", "createdDate"},
			JsonPath:   contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		listFirewallNameFilter = viper.GetString("name")

		viper.BindPFlag("firewallId", cmd.Flags().Lookup("firewallId"))
		listFirewallIdFilter = viper.GetString("firewallId")

		viper.BindPFlag("status", cmd.Flags().Lookup("status"))
		listFirewallStatusFilter = viper.GetString("status")

		viper.BindPFlag("instances", cmd.Flags().Lookup("instances"))
		listFirewallInstancesFilter = viper.GetString("instances")

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(firewallsGetCmd)

	firewallsGetCmd.Flags().StringVarP(&listFirewallNameFilter, "name", "n", "",
		`Filter by firewall name`)

	firewallsGetCmd.Flags().StringVar(&listFirewallIdFilter, "firewallId", "",
		`Filter by firewall ID`)

	firewallsGetCmd.Flags().StringVar(&listFirewallStatusFilter, "status", "",
		`Filter by firewall status`)

	firewallsGetCmd.Flags().StringVar(&listFirewallInstancesFilter, "instances", "",
		`Filter by instance IDs (comma-separated), e.g. --instances "101010,101001,202002"`)
}
