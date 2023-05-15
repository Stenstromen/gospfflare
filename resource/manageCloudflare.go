package resource

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func manageCloudflare(nameanddomain string, recordtype string, method string, requestBody string) error {
	url := "https://api.cloudflare.com/client/v4/zones"
	var bearer = "Bearer " + os.Getenv("TOKEN")

	initialRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	initialRequest.Header.Add("Authorization", bearer)
	client := &http.Client{}
	initialResponse, err := client.Do(initialRequest)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return err
	}
	defer initialResponse.Body.Close()

	body, err := io.ReadAll(initialResponse.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return err
	}

	var res Res
	if err := json.Unmarshal(body, &res); err != nil {
		log.Println(err)
		return err
	}

	var did string
	searchurl := "https://api.cloudflare.com/client/v4/zones/" + res.Result[0].ID + "/dns_records"
	searchRequest, err := http.NewRequest("GET", searchurl, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	searchRequest.Header.Add("Authorization", bearer)
	searchResponse, err := client.Do(searchRequest)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return err
	}
	defer searchResponse.Body.Close()

	searchBody, err := io.ReadAll(searchResponse.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return err
	}

	var recordsres RecordsRes
	if err := json.Unmarshal(searchBody, &recordsres); err != nil {
		log.Println(err)
		return err
	}

	for i := range recordsres.Result {
		if nameanddomain == recordsres.Result[i].Name {
			did = recordsres.Result[i].ID
			break
		}
	}

	var manageurl string
	if method == "PUT" {
		manageurl = "https://api.cloudflare.com/client/v4/zones/" + res.Result[0].ID + "/dns_records/" + did
	} else { // POST
		manageurl = "https://api.cloudflare.com/client/v4/zones/" + res.Result[0].ID + "/dns_records"
	}

	var jsonStr = []byte(requestBody)
	manageRequest, err := http.NewRequest(method, manageurl, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
		return err
	}

	manageRequest.Header.Set("Content-Type", "application/json")
	manageRequest.Header.Add("Authorization", bearer)

	manageResponse, err := client.Do(manageRequest)
	if err != nil {
		log.Println(err)
		return err
	}
	defer manageResponse.Body.Close()

	log.Println("Cloudflare Response Status for: "+recordtype, manageResponse.Status)
	return nil
}
