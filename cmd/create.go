package cmd

import (
	"gospfflare/resource"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new SPF record",
	Long:  `Create a new SPF record`,
	RunE:  resource.ResourceCreate,
}

func init() {
	var ip4s []string
	var ip6s []string
	var includes []string
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("domain", "d", "", "Domain name")
	createCmd.Flags().BoolP("a", "a", false, "Will match if the domain name has an address record (Ipv4/6) that can be resolved to the sender's address")
	createCmd.Flags().BoolP("mx", "m", false, "Will match if the domain name has a mail exchanger record (MX) resolving to the sender's address")
	createCmd.Flags().StringArrayVarP(&ip4s, "ip4", "4", []string{}, "IPv4 address")
	createCmd.Flags().StringArrayVarP(&ip6s, "ip6", "6", []string{}, "IPv6 address")
	createCmd.Flags().StringArrayVarP(&includes, "include", "i", []string{}, "Include SPF record")
	createCmd.Flags().BoolP("neutral", "n", false, "NEUTRAL policy")
	createCmd.Flags().BoolP("softfail", "s", false, "SOFTFAIL policy")
	createCmd.Flags().BoolP("fail", "f", false, "FAIL policy")
	createCmd.MarkFlagRequired("domain")
}
