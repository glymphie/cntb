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

var firewallUpdateCmd = &cobra.Command{
	Use:   "firewall [firewallId]",
	Short: "Updates a specific firewall",
	Long:  `Updates the specific firewall by setting new values either by file input or flags / environment variables`,
	Run: func(cmd *cobra.Command, args []string) {
		updatefilewallRequest := *firewallClient.NewPatchFirewallRequestWithDefaults()

		content := contaboCmd.OpenStdinOrFile()
		switch content {
		case nil:
			// from arguments
			if updateFirewallName != "" {
				updatefilewallRequest.Name = &updateFirewallName
			}
			if updateFirewallStatus != "" {
				updatefilewallRequest.Status = &updateFirewallStatus
			}
			if updateFirewallDescription != "" {
				updatefilewallRequest.Description = &updateFirewallDescription
			}
		default:
			// from file / stdin
			var requestFromFile firewallClient.PatchFirewallRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge updatefilewallRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&updatefilewallRequest)
		}

		resp, httpResp, err := client.ApiClient().FirewallsApi.
			PatchFirewall(context.Background(), updateFirewallId).PatchFirewallRequest(updatefilewallRequest).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while updating firewall")

		responseJSON, _ := resp.MarshalJSON()
		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide a firewallId.")
		}

		updateFirewallId = args[0]
		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		updateFirewallName = viper.GetString("name")
		viper.BindPFlag("status", cmd.Flags().Lookup("status"))
		updateFirewallStatus = viper.GetString("status")
		viper.BindPFlag("description", cmd.Flags().Lookup("description"))
		updateFirewallDescription = viper.GetString("description")

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(firewallUpdateCmd)

	firewallUpdateCmd.Flags().StringVarP(&updateFirewallName, "name", "n", "",
		`name of the firwall`)

	firewallUpdateCmd.Flags().StringVarP(&updateFirewallStatus, "status", "s", "",
		`status of the firewall`)

	firewallUpdateCmd.Flags().StringVarP(&updateFirewallDescription, "description", "", "",
		`description of the firewall`)
}
