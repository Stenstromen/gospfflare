package cmd

import (
	"gospfflare/resource"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing SPF record",
	Long:  `Update an existing SPF record. Will overwrite existing record.`,
	RunE:  resource.ResourceUpdate,
}

func init() {
	var ip4s []string
	var ip6s []string
	var includes []string
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("domain", "d", "", "Domain name")
	updateCmd.Flags().BoolP("a", "a", false, "Will match if the domain name has an address record (Ipv4/6) that can be resolved to the sender's address")
	updateCmd.Flags().BoolP("mx", "m", false, "Will match if the domain name has a mail exchanger record (MX) resolving to the sender's address")
	updateCmd.Flags().StringArrayVarP(&ip4s, "ip4", "4", []string{}, "IPv4 address")
	updateCmd.Flags().StringArrayVarP(&ip6s, "ip6", "6", []string{}, "IPv6 address")
	updateCmd.Flags().StringArrayVarP(&includes, "include", "i", []string{}, "Include SPF record")
	updateCmd.Flags().BoolP("neutral", "n", false, "NEUTRAL policy")
	updateCmd.Flags().BoolP("softfail", "s", false, "SOFTFAIL policy")
	updateCmd.Flags().BoolP("fail", "f", false, "FAIL policy")
	updateCmd.MarkFlagRequired("domain")
}
