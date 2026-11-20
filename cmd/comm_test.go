package cmd

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	alidns "github.com/alibabacloud-go/alidns-20150109/v5/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

func TestDNSRequests(t *testing.T) {
	for _, tc := range []struct {
		action string
		fields map[string]string
		call   func() error
	}{
		{"DescribeDomainRecords", map[string]string{"DomainName": ServiceDomain, "PageSize": "500"}, func() error {
			records, err := describeRecords(ServiceDomain)
			if err == nil && (len(records) != 1 || records[0] != (dnsRecord{RecordId: "123", Type: "SRV", RR: "_test._tcp", Value: "10 20 443 example.com", Remark: "test"})) {
				t.Errorf("unexpected records: %+v", records)
			}
			return err
		}},
		{"AddDomainRecord", map[string]string{"DomainName": ServiceDomain, "Type": "SRV", "RR": "_test._tcp", "Value": "10 20 443 example.com"}, func() error { return addRecord(ServiceDomain, "SRV", "_test._tcp", "10 20 443 example.com") }},
		{"UpdateDomainRecord", map[string]string{"RecordId": "123", "Type": "SRV", "RR": "_test._tcp", "Value": "10 20 443 example.com"}, func() error { return updateRecord("123", "SRV", "_test._tcp", "10 20 443 example.com") }},
		{"UpdateDomainRecordRemark", map[string]string{"RecordId": "123", "Remark": "test"}, func() error { return updateRecordRemark("123", "test") }},
		{"DeleteDomainRecord", map[string]string{"RecordId": "123"}, func() error { return deleteRecord("123") }},
	} {
		t.Run(tc.action, func(t *testing.T) {
			calls := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				if r.Header.Get("x-acs-action") != tc.action {
					t.Errorf("action = %q", r.Header.Get("x-acs-action"))
				}
				if err := r.ParseForm(); err != nil {
					t.Error(err)
				}
				for key, want := range tc.fields {
					if got := r.Form.Get(key); got != want {
						t.Errorf("%s = %q, want %q", key, got, want)
					}
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"RequestId":"request-test","RecordId":"123","DomainRecords":{"Record":[{"RecordId":"123","Type":"SRV","RR":"_test._tcp","Value":"10 20 443 example.com","Remark":"test"}]}}`))
			}))
			defer server.Close()
			useTestDNSClient(t, server.URL)
			if err := tc.call(); err != nil {
				t.Fatal(err)
			}
			if calls != 1 {
				t.Errorf("requests = %d", calls)
			}
		})
	}
}

func useTestDNSClient(t *testing.T, endpoint string) {
	t.Helper()
	old := newDNSClient
	t.Cleanup(func() { newDNSClient = old })
	newDNSClient = func() (*alidns.Client, error) {
		return alidns.NewClient(&openapi.Config{
			AccessKeyId: dara.String("test-id"), AccessKeySecret: dara.String("test-secret"),
			Endpoint: dara.String(strings.TrimPrefix(endpoint, "http://")), Protocol: dara.String("HTTP"),
		})
	}
}

func TestDNSPermissionError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"Code":"Forbidden.RAM","Message":"not authorized","RequestId":"denied-request"}`))
	}))
	defer server.Close()
	useTestDNSClient(t, server.URL)
	_, err := describeRecords(ServiceDomain)
	if err == nil {
		t.Fatal("expected permission error")
	}
	for _, want := range []string{"Forbidden.RAM", "denied-request"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q missing %q", err, want)
		}
	}
}
