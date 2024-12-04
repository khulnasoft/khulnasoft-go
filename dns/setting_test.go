// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/dns"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
)

func TestSettingEditWithOptionalParams(t *testing.T) {
	t.Skip("HTTP 422 from prism")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := khulnasoft.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.DNS.Settings.Edit(context.TODO(), dns.SettingEditParams{
		AccountID: khulnasoft.F("account_id"),
		ZoneDefaults: khulnasoft.F(dns.DNSSettingParam{
			FlattenAllCNAMEs: khulnasoft.F(false),
			FoundationDNS:    khulnasoft.F(false),
			MultiProvider:    khulnasoft.F(false),
			Nameservers: khulnasoft.F(dns.NameserverParam{
				Type: khulnasoft.F(dns.NameserverTypeCloudflareStandard),
			}),
			NSTTL:              khulnasoft.F(86400.000000),
			SecondaryOverrides: khulnasoft.F(false),
			SOA: khulnasoft.F(dns.DNSSettingSOAParam{
				Expire:  khulnasoft.F(604800.000000),
				MinTTL:  khulnasoft.F(1800.000000),
				MNAME:   khulnasoft.F("kristina.ns.khulnasoft.com"),
				Refresh: khulnasoft.F(10000.000000),
				Retry:   khulnasoft.F(2400.000000),
				RNAME:   khulnasoft.F("admin.example.com"),
				TTL:     khulnasoft.F(3600.000000),
			}),
			ZoneMode: khulnasoft.F(dns.DNSSettingZoneModeStandard),
		}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingGetWithOptionalParams(t *testing.T) {
	t.Skip("HTTP 422 from prism")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := khulnasoft.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.DNS.Settings.Get(context.TODO(), dns.SettingGetParams{
		AccountID: khulnasoft.F("account_id"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
