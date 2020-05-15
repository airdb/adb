package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var hostCommand = &cobra.Command{
	Use:                "host",
	Short:              "host operation",
	Long:               "host operation",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		host()
	},
}

type HostInstance struct {
	PageNumber    int    `json:"PageNumber"`
	PageSize      int    `json:"PageSize"`
	RequestID     string `json:"RequestId"`
	TotalCount    int    `json:"TotalCount"`
	DomainRecords struct {
		Record []struct {
			DomainName string `json:"DomainName"`
			Line       string `json:"Line"`
			Locked     bool   `json:"Locked"`
			RR         string `json:"RR"`
			RecordID   string `json:"RecordId"`
			Status     string `json:"Status"`
			TTL        int    `json:"TTL"`
			Type       string `json:"Type"`
			Value      string `json:"Value"`
			Weight     int    `json:"Weight"`
		} `json:"Record"`
	} `json:"DomainRecords"`
}

func host() {
	// alidns DescribeDomainRecords --DomainName airdb.host
	domain := "airdb.host"
	cmd := exec.Command("aliyun", "alidns", "DescribeDomainRecords", "--DomainName", domain)

	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	var instance HostInstance

	err = json.Unmarshal(out.Bytes(), &instance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-8s\t%-32s\t%s\n", "Type", "Host", "IP")

	typ := "host"

	for _, record := range instance.DomainRecords.Record {
		if strings.HasPrefix(record.Value, ".docker") {
			typ = "docker"
		}

		if record.Type == "A" {
			fmt.Printf("%-8s%-32s\t%s\n", typ, record.RR+"."+record.DomainName, record.Value)
		}
	}
}
