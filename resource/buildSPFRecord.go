package resource

func buildSPFRecord(a bool, mx bool, ip4 []string, ip6 []string, include []string, policy string) string {
	spf := "v=spf1 "

	if a {
		spf += "a "
	}

	if mx {
		spf += "mx "
	}

	for _, ip4s := range ip4 {
		spf += "ip4:" + ip4s + " "
	}

	for _, ip6s := range ip6 {
		spf += "ip6:" + ip6s + " "
	}

	for _, includes := range include {
		spf += "include:" + includes + " "
	}

	if policy == "" {
		spf += "-all"
	} else {
		spf += policy + "all"
	}

	return spf
}
