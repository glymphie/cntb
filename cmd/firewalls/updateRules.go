package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	firewallClient "contabo.com/cli/cntb/openapi"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ValidRules struct {
	Inbound *[]firewallClient.RulesRequest `json:"inbound"`
}

var firewallUpdateRulesCmd = &cobra.Command{
	Use:     "firewall-rules [firewallId]",
	Short:   "Update the rules of a firewall by firewallId.",
	Long:    `Update the rules of a firewall by firewallId based on a json/yaml that is provided.`,
	Example: `cntb update firewall-rules 7581640c-fd89-4f7a-9f2e-3b1ab135ae0d --rules='{...}'`,
	Run: func(cmd *cobra.Command, args []string) {
		updateRulefForFirewallRequst := *firewallClient.NewPutFirewallRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			if addRulesFirewallRules != "" {
				var checkRules *ValidRules
				var rulesRequest firewallClient.RulesRequest

				err := json.Unmarshal([]byte(addRulesFirewallRules), &checkRules)
				if err != nil {
					log.Error("Invalid `rules`. Please check the JSON syntax.")
					log.Fatal(fmt.Sprintf("Error: %v", err))
				}

				if checkRules.Inbound == nil && addRulesFirewallRules != "{}" {
					log.Error("Error while updating rules: 400 - Bad Request only inbound rules are allowed!")
					log.Fatal("Aborting, due to error")
				} else {
					json.Unmarshal([]byte(addRulesFirewallRules), &rulesRequest)
				}
				updateRulefForFirewallRequst.Rules = &rulesRequest
			}

		default:
			// from file / stdin
			var requestFromFile firewallClient.PutFirewallRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge addRulesToFirewallRequst with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&updateRulefForFirewallRequst)
		}

		resp, httpResp, err := client.ApiClient().FirewallsApi.
			PutFirewall(context.Background(), addRulesFirewallId).
			XRequestId(uuid.NewV4().String()).
			PutFirewallRequest(updateRulefForFirewallRequst).Execute()

		util.HandleErrors(err, httpResp, "while adding rules to firewall")
		responseJson, _ := json.Marshal(resp.Data)

		if viper.GetString("output") != "json" && viper.GetString("output") != "yaml" {
			firewalls := make([]jmap, 0)
			for _, firewall := range resp.Data {
				updatedFirewall, _ := util.StructToMap(firewall)
				var instances []int64
				for _, instance := range firewall.Instances {
					instances = append(instances, instance.InstanceId)
				}
				if instances == nil {
					updatedFirewall["instances"] = "[]"
				} else {
					updatedFirewall["instances"] = instances
				}
				if firewall.Rules.Inbound == nil {
					updatedFirewall["rules"] = "{[]}"
				}

				firewalls = append(firewalls, updatedFirewall)

			}
			responseJson, _ = json.Marshal(firewalls)
		}

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"firewallId", "name", "status", "isDefault", "instances", "rules",
			},
			WideFilter: []string{
				"firewallId", "name", "status", "isDefault", "description", "instances", "createdDate", "rules",
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
			log.Fatal("Please provide a firewallId.")
		}

		addRulesFirewallId = args[0]

		viper.BindPFlag("rules", cmd.Flags().Lookup("rules"))
		addRulesFirewallRules = viper.GetString("rules")

		if addRulesFirewallId == "" {
			cmd.Help()
			log.Fatal("Argument firewallId is empty. Please provide a non empty firewallId.")
		}

		if addRulesFirewallRules == "" {
			cmd.Help()
			log.Fatal("Argument addRulesFirewallRules is empty. Please provide non empty addRulesFirewallRules.")
		}
		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(firewallUpdateRulesCmd)
	firewallUpdateRulesCmd.Flags().StringVarP(&addRulesFirewallRules, "rules", "r", "",
		"firewall rules")
}
