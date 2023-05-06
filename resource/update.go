package resource

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

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

	putToCloudflare(domain, domain+" TXT", genCloudflareReq("TXT", "@", buildSPFRecord(a, mx, ip4, ip6, include, policy), "Updated"))

	return nil
}

func putToCloudflare(nameanddomain string, recordtype string, putBody string) {
	url := "https://api.cloudflare.com/client/v4/zones"
	var bearer = "Bearer " + os.Getenv("TOKEN")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var res Res
	if err := json.Unmarshal(body, &res); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	searchurl := "https://api.cloudflare.com/client/v4/zones/" + res.Result[0].ID + "/dns_records"
	req2, err2 := http.NewRequest("GET", searchurl, nil)
	if err2 != nil {
		log.Println(err2)
		os.Exit(1)
	}

	req2.Header.Add("Authorization", bearer)
	client2 := &http.Client{}
	resp2, err2 := client2.Do(req2)
	if err2 != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp2.Body.Close()
	body2, err2 := io.ReadAll(resp2.Body)
	if err2 != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var recordsres RecordsRes
	if err2 := json.Unmarshal(body2, &recordsres); err2 != nil {
		log.Println(err2)
		os.Exit(1)
	}

	var did string

	for i := range recordsres.Result {
		if nameanddomain == recordsres.Result[i].Name {
			did = recordsres.Result[i].ID
		}
	}

	puturl := "https://api.cloudflare.com/client/v4/zones/" + res.Result[0].ID + "/dns_records/" + did

	var jsonStr = []byte(putBody)
	req3, err3 := http.NewRequest("PATCH", puturl, bytes.NewBuffer(jsonStr))
	if err3 != nil {
		log.Println(err3)
		os.Exit(1)
	}

	req3.Header.Set("Content-Type", "application/json")
	req3.Header.Add("Authorization", bearer)

	client3 := &http.Client{}
	resp3, err3 := client3.Do(req3)
	if err3 != nil {
		log.Println(err3)
		os.Exit(1)
	}
	defer resp.Body.Close()

	log.Println("Cloudflare Response Status for: "+recordtype, resp3.Status)
}
