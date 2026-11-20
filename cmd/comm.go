package cmd

import (
	"fmt"

	"github.com/airdb/adb/internal/adblib"
	alidns "github.com/alibabacloud-go/alidns-20150109/v5/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

const (
	ServiceDomain = "airdb.space"
	HostDomain    = "airdb.host"
)

const describePageSize = 500

// dnsRecord keeps SDK pointer fields out of command output and filtering.
type dnsRecord struct {
	RecordId string
	Type     string
	RR       string
	Value    string
	Remark   string
}

func aliyunConfigInit() (*alidns.Client, error) {
	return alidns.NewClient(&openapi.Config{
		RegionId:        dara.String(adblib.AdbConfig.AliyunRegionID),
		AccessKeyId:     dara.String(adblib.AdbConfig.AliyunAccessKeyID),
		AccessKeySecret: dara.String(adblib.AdbConfig.AliyunAccessKeySecret),
		Endpoint:        dara.String("alidns.aliyuncs.com"),
		Protocol:        dara.String("HTTPS"),
	})
}

var newDNSClient = aliyunConfigInit

func describeRecords(domain string) ([]dnsRecord, error) {
	client, err := newDNSClient()
	if err != nil {
		return nil, err
	}

	request := &alidns.DescribeDomainRecordsRequest{}
	request.DomainName = dara.String(domain)
	request.PageSize = dara.Int64(describePageSize)

	output, err := client.DescribeDomainRecords(request)
	if err != nil {
		return nil, err
	}

	records := make([]dnsRecord, 0)
	if output == nil || output.Body == nil || output.Body.DomainRecords == nil {
		return records, nil
	}
	for _, record := range output.Body.DomainRecords.Record {
		if record == nil {
			continue
		}
		records = append(records, dnsRecord{
			RecordId: dara.StringValue(record.RecordId),
			Type:     dara.StringValue(record.Type),
			RR:       dara.StringValue(record.RR),
			Value:    dara.StringValue(record.Value),
			Remark:   dara.StringValue(record.Remark),
		})
	}
	return records, nil
}

func addRecord(domain, recordType, rr, value string) error {
	client, err := newDNSClient()
	if err != nil {
		return err
	}

	request := &alidns.AddDomainRecordRequest{}
	request.DomainName = dara.String(domain)
	request.Type = dara.String(recordType)
	request.RR = dara.String(rr)
	request.Value = dara.String(value)

	output, err := client.AddDomainRecord(request)
	if err != nil {
		return err
	}

	fmt.Println(output)

	return nil
}

func updateRecord(recordID, recordType, rr, value string) error {
	client, err := newDNSClient()
	if err != nil {
		return err
	}

	request := &alidns.UpdateDomainRecordRequest{}
	request.RecordId = dara.String(recordID)
	request.Type = dara.String(recordType)
	request.RR = dara.String(rr)
	request.Value = dara.String(value)

	output, err := client.UpdateDomainRecord(request)
	if err != nil {
		return err
	}

	fmt.Println(output)

	return nil
}

func updateRecordRemark(recordID, remark string) error {
	client, err := newDNSClient()
	if err != nil {
		return err
	}

	request := &alidns.UpdateDomainRecordRemarkRequest{}
	request.RecordId = dara.String(recordID)
	request.Remark = dara.String(remark)

	output, err := client.UpdateDomainRecordRemark(request)
	if err != nil {
		return err
	}

	fmt.Println(output)

	return nil
}

func deleteRecord(recordID string) error {
	client, err := newDNSClient()
	if err != nil {
		return err
	}

	request := &alidns.DeleteDomainRecordRequest{}
	request.RecordId = dara.String(recordID)

	output, err := client.DeleteDomainRecord(request)
	if err != nil {
		return err
	}

	fmt.Println(output)

	return nil
}
