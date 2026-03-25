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
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var firewallCreateCmd = &cobra.Command{
	Use:   "firewall",
	Short: "Creates a new firewall.",
	Long:  `Creates a new firewall based on a json/yaml that is provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		createFirewallRequest := *firewallClient.NewCreateFirewallRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()
		switch content {
		case nil:
			// from arguments
			createFirewallRequest.Name = createFirewallName
			createFirewallRequest.Status = createFirewallStatus
			if createFirewallDescription != "" {
				createFirewallRequest.Description = &createFirewallDescription
			}
			if createFirewallRules != "" {
				var rulesRequest firewallClient.RulesRequest
				err := json.Unmarshal([]byte(createFirewallRules), &rulesRequest)
				if err != nil {
					log.Error("Invalid `rules`. Please check the JSON syntax.")
					log.Fatal(fmt.Sprintf("Error: %v", err))
				}
				createFirewallRequest.Rules = &rulesRequest
			}
		default:
			// from file / stdin
			var requestFromFile firewallClient.CreateFirewallRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge createFirewallRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&createFirewallRequest)
		}

		resp, httpResp, err := client.ApiClient().FirewallsApi.CreateFirewall(context.Background()).
			XRequestId(uuid.NewV4().String()).CreateFirewallRequest(createFirewallRequest).Execute()

		util.HandleErrors(err, httpResp, "while creating firewall")

		fmt.Printf("%v\n", resp.Data[0].FirewallId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		createFirewallName = viper.GetString("name")

		viper.BindPFlag("rules", cmd.Flags().Lookup("rules"))
		createFirewallRules = viper.GetString("rules")

		viper.BindPFlag("status", cmd.Flags().Lookup("status"))
		createFirewallStatus = viper.GetString("status")

		viper.BindPFlag("description", cmd.Flags().Lookup("description"))
		createFirewallDescription = viper.GetString("description")

		if contaboCmd.InputFile == "" {
			// arguments required
			if createFirewallName == "" {
				cmd.Help()
				log.Fatal("Argument name is empty. Please provide one.")
			}
			if createFirewallStatus == "" {
				cmd.Help()
				log.Fatal("Argument status is empty. Please provide one.")
			}
		}
		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(firewallCreateCmd)

	firewallCreateCmd.Flags().StringVarP(&createFirewallName, "name", "n", "", `name of the firewall`)

	firewallCreateCmd.Flags().StringVarP(&createFirewallStatus, "status", "s", "",
		"firewall status. Can be active or inactive")

	firewallCreateCmd.Flags().StringVarP(&createFirewallDescription, "description", "", "",
		"firewall description")

	firewallCreateCmd.Flags().StringVarP(&createFirewallRules, "rules", "r", "",
		"firewall rules")
}
