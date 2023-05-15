package resource

import (
	"github.com/spf13/cobra"
)

func ResourceUpdate(cmd *cobra.Command, args []string) error {
	var policy string
	domain, err := cmd.Flags().GetString("domain")
	if err != nil {
		return err
	}

	a, err := cmd.Flags().GetBool("a")
	if err != nil {
		return err
	}

	mx, err := cmd.Flags().GetBool("mx")
	if err != nil {
		return err
	}

	ip4, err := cmd.Flags().GetStringArray("ip4")
	if err != nil {
		return err
	}

	ip6, err := cmd.Flags().GetStringArray("ip6")
	if err != nil {
		return err
	}

	include, err := cmd.Flags().GetStringArray("include")
	if err != nil {
		return err
	}

	neutral, err := cmd.Flags().GetBool("neutral")
	if neutral {
		policy = "?"
	}
	if err != nil {
		return err
	}

	softfail, err := cmd.Flags().GetBool("softfail")
	if softfail {
		policy = "~"
	}
	if err != nil {
		return err
	}

	fail, err := cmd.Flags().GetBool("fail")
	if fail {
		policy = "-"
	}
	if err != nil {
		return err
	}

	manageCloudflare(domain, "TXT", "PUT", genCloudflareReq("TXT", "@", buildSPFRecord(a, mx, ip4, ip6, include, policy), "Updated"))

	return nil
}
