package main

import (
	"fmt"
	"log"
	"os"

	"github.com/j4ng5y/cfupdater/config"
	cf "github.com/j4ng5y/cfupdater/updater/cloudflare"
	ip "github.com/j4ng5y/cfupdater/updater/ipinfo"
	"github.com/spf13/cobra"
)

var (
	configFilePath string

	// CFupdaterCmd is the main command of cfupdater
	CFupdaterCmd = &cobra.Command{
		Use:     "cfupdater",
		Version: "0.2.1",
		Short:   "cfupdater is an app to update cloudflare DNS resouce records with the contents of ipinfo.io",
		Run:     cfupdaterFunc,
	}

	cloudflareAPIToken   string
	cloudflareZoneID     string
	cloudflareRecordName string
	ipinfoAPIToken       string

	// CFupdaterConfigCmd is the configuration command of cfupdater
	CFupdaterConfigCmd = &cobra.Command{
		Use:   "configure",
		Short: "configure generates a configuration file",
		Run:   cfupdaterConfigFunc,
	}
)

func cfupdaterFunc(ccmd *cobra.Command, args []string) {
	config, err := config.New(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	recReq := &cf.DNSRecordRequest{
		URL:        fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records?name=%s", config.CloudFlare.DNSRecord.ZoneID, config.CloudFlare.DNSRecord.RecordName),
		APIToken:   config.CloudFlare.General.APIToken,
		ZoneID:     config.CloudFlare.DNSRecord.ZoneID,
		RecordName: config.CloudFlare.DNSRecord.RecordName,
	}
	recResp, err := recReq.Get()
	if err != nil {
		log.Fatal(err)
	}

	ipReq := &ip.IPInfoRequest{
		URL:      fmt.Sprintf("https://ipinfo.io"),
		APIToken: config.IPInfo.General.APIToken,
	}
	ipResp, err := ipReq.Get()
	if err != nil {
		log.Fatal(err)
	}

	if recResp.Result[0].Content == ipResp.IP {
		log.Printf("IPInfo IP Address: %s, CloudFlare DNS Record Address: %s, Result: No Update Needed\n", ipResp.IP, recResp.Result[0].Content)
		os.Exit(0)
	}

	recUpdateReq := &cf.UpdateDNSRecordRequest{
		URL:        fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", recResp.Result[0].ZoneID, recResp.Result[0].ID),
		APIToken:   config.CloudFlare.General.APIToken,
		ZoneID:     config.CloudFlare.DNSRecord.ZoneID,
		RecordName: config.CloudFlare.DNSRecord.RecordName,
		Type:       recResp.Result[0].Type,
		Name:       recResp.Result[0].Name,
		Content:    ipResp.IP,
		TTL:        recResp.Result[0].TTL,
		Proxied:    recResp.Result[0].Proxied,
	}
	recUpdateResp, err := recUpdateReq.Update()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("IPInfo IP Address: %s, CloudFlare DNS Record Address: %s, Result: %v\n", ipResp.IP, recResp.Result[0].Content, recUpdateResp)
}

func cfupdaterConfigFunc(ccmd *cobra.Command, args []string) {
	config := &config.Config{}
	config.CloudFlare.General.APIToken = cloudflareAPIToken
	config.CloudFlare.DNSRecord.ZoneID = cloudflareZoneID
	config.CloudFlare.DNSRecord.RecordName = cloudflareRecordName
	config.IPInfo.General.APIToken = ipinfoAPIToken

	if err := config.Write(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	CFupdaterCmd.Flags().StringVarP(&configFilePath, "config-file", "c", "", "The path to the config file to use")
	CFupdaterCmd.MarkFlagRequired("config-file")

	CFupdaterCmd.AddCommand(CFupdaterConfigCmd)

	CFupdaterConfigCmd.Flags().StringVar(&cloudflareAPIToken, "cloudflare-api-token", "", "The Cloudflare API token to use")
	CFupdaterConfigCmd.Flags().StringVar(&cloudflareZoneID, "cloudflare-dns-zone-id", "", "The Cloudflare DNS Zone ID to use")
	CFupdaterConfigCmd.Flags().StringVar(&cloudflareRecordName, "cloudflare-dns-record-name", "", "The Cloudflare DNS Record Name to update")
	CFupdaterConfigCmd.Flags().StringVar(&ipinfoAPIToken, "ipinfo-api-token", "", "The IPInfo API token to use")

	CFupdaterConfigCmd.MarkFlagRequired("cloudflare-api-token")
	CFupdaterConfigCmd.MarkFlagRequired("cloudflare-dns-zone-id")
	CFupdaterConfigCmd.MarkFlagRequired("cloudflare-dns-record-name")
	CFupdaterConfigCmd.MarkFlagRequired("ipinfo-api-token")
}

func main() {
	if err := CFupdaterCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
