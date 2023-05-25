package adblib

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

const HostDoamin = "airdb.host"

func GetHosts() {
	client, err := NewAliyunClient()
	if err != nil {
		panic(err)
	}

	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = HostDoamin

	output, err := client.DescribeDomainRecords(request)
	if err != nil {
		fmt.Println(err)
	}

	for _, rr := range output.DomainRecords.Record {
		if rr.Type != "A" {
			continue
		}

		/*
			if rr.RR == sailor.DelimiterStar || rr.RR == sailor.DelimiterAt {
				continue
			}
		*/

		// fmt.Printf("%-20s\t%-32s\t%s\n", rr.RecordId, rr.RR, rr.Value)
		fmt.Printf("%-20s %-32s %-64s %s\n", rr.RecordId, rr.RR, rr.Value, rr.Remark)
	}
}
