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

var dnsGetCmd = &cobra.Command{
	Use:   "dns [zone]",
	Short: "list DNS records for a zone",
	Long:  "Retrieves DNS records for a DNS zone",
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().
			DNSApi.RetrieveDnsZoneRecordsList(context.Background(), getZoneName).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving DNS zone records")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"name", "type", "data", "ttl"},
			WideFilter: []string{"name", "type", "data", "ttl", "prio"},
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
		if len(args) != 1 {
			cmd.Help()
			log.Fatal("Please provide exactly one zone name.")
		}
		getZoneName = args[0]

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(dnsGetCmd)
}
