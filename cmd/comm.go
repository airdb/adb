package cmd

import (
	"fmt"

	"github.com/airdb/adb/internal/adblib"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

const (
	ServiceDomain = "airdb.space"
	HostDomain    = "airdb.host"
)

const describePageSize = 500

func aliyunConfigInit() (*alidns.Client, error) {
	return alidns.NewClientWithAccessKey(
		adblib.AdbConfig.AliyunRegionID,
		adblib.AdbConfig.AliyunAccessKeyID,
		adblib.AdbConfig.AliyunAccessKeySecret,
	)
}

func describeRecords(domain string) ([]alidns.Record, error) {
	client, err := aliyunConfigInit()
	if err != nil {
		return nil, err
	}

	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = domain
	request.PageSize = requests.NewInteger(describePageSize)

	output, err := client.DescribeDomainRecords(request)
	if err != nil {
		return nil, err
	}

	return output.DomainRecords.Record, nil
}

func addRecord(domain, recordType, rr, value string) error {
	client, err := aliyunConfigInit()
	if err != nil {
		return err
	}

	request := alidns.CreateAddDomainRecordRequest()
	request.DomainName = domain
	request.Type = recordType
	request.RR = rr
	request.Value = value

	output, err := client.AddDomainRecord(request)
	if err != nil {
		return err
	}

	fmt.Println(output)

	return nil
}

func updateRecord(recordID, recordType, rr, value string) error {
	client, err := aliyunConfigInit()
	if err != nil {
		return err
	}

	request := alidns.CreateUpdateDomainRecordRequest()
	request.RecordId = recordID
	request.Type = recordType
	request.RR = rr
	request.Value = value

	output, err := client.UpdateDomainRecord(request)
	if err != nil {
		return err
	}

	fmt.Println(output)

	return nil
}

func updateRecordRemark(recordID, remark string) error {
	client, err := aliyunConfigInit()
	if err != nil {
		return err
	}

	request := alidns.CreateUpdateDomainRecordRemarkRequest()
	request.RecordId = recordID
	request.Remark = remark

	output, err := client.UpdateDomainRecordRemark(request)
	if err != nil {
		return err
	}

	fmt.Println(output)

	return nil
}

func deleteRecord(recordID string) error {
	client, err := aliyunConfigInit()
	if err != nil {
		return err
	}

	request := alidns.CreateDeleteDomainRecordRequest()
	request.RecordId = recordID

	output, err := client.DeleteDomainRecord(request)
	if err != nil {
		return err
	}

	fmt.Println(output)

	return nil
}
