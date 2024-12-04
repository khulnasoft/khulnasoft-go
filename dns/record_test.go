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
	"github.com/khulnasoft/khulnasoft-go/shared"
)

func TestRecordNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.DNS.Records.New(context.TODO(), dns.RecordNewParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Record: dns.RecordParam{
			Name:     khulnasoft.F("example.com"),
			Comment:  khulnasoft.F("Domain verification record"),
			Proxied:  khulnasoft.F(true),
			Settings: khulnasoft.F[any](map[string]interface{}{}),
			Tags:     khulnasoft.F([]dns.RecordTagsParam{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
			TTL:      khulnasoft.F(dns.TTL1),
		},
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecordUpdateWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.DNS.Records.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.RecordUpdateParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Record: dns.RecordParam{
				Name:     khulnasoft.F("example.com"),
				Comment:  khulnasoft.F("Domain verification record"),
				Proxied:  khulnasoft.F(true),
				Settings: khulnasoft.F[any](map[string]interface{}{}),
				Tags:     khulnasoft.F([]dns.RecordTagsParam{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
				TTL:      khulnasoft.F(dns.TTL1),
			},
		},
	)
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecordListWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Records.List(context.TODO(), dns.RecordListParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Comment: khulnasoft.F(dns.RecordListParamsComment{
			Absent:     khulnasoft.F("absent"),
			Contains:   khulnasoft.F("ello, worl"),
			Endswith:   khulnasoft.F("o, world"),
			Exact:      khulnasoft.F("Hello, world"),
			Present:    khulnasoft.F("present"),
			Startswith: khulnasoft.F("Hello, w"),
		}),
		Content:   khulnasoft.F("127.0.0.1"),
		Direction: khulnasoft.F(shared.SortDirectionAsc),
		Match:     khulnasoft.F(dns.RecordListParamsMatchAny),
		Name:      khulnasoft.F("example.com"),
		Order:     khulnasoft.F(dns.RecordListParamsOrderType),
		Page:      khulnasoft.F(1.000000),
		PerPage:   khulnasoft.F(5.000000),
		Proxied:   khulnasoft.F(true),
		Search:    khulnasoft.F("www.khulnasoft.com"),
		Tag: khulnasoft.F(dns.RecordListParamsTag{
			Absent:     khulnasoft.F("important"),
			Contains:   khulnasoft.F("greeting:ello, worl"),
			Endswith:   khulnasoft.F("greeting:o, world"),
			Exact:      khulnasoft.F("greeting:Hello, world"),
			Present:    khulnasoft.F("important"),
			Startswith: khulnasoft.F("greeting:Hello, w"),
		}),
		TagMatch: khulnasoft.F(dns.RecordListParamsTagMatchAny),
		Type:     khulnasoft.F(dns.RecordListParamsTypeA),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecordDelete(t *testing.T) {
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
	_, err := client.DNS.Records.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.RecordDeleteParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		},
	)
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecordEditWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.DNS.Records.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.RecordEditParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Record: dns.RecordParam{
				Comment:  khulnasoft.F("Domain verification record"),
				Name:     khulnasoft.F("example.com"),
				Proxied:  khulnasoft.F(true),
				Settings: khulnasoft.F[any](map[string]interface{}{}),
				Tags:     khulnasoft.F([]dns.RecordTagsParam{"owner:dns-team", "owner:dns-team", "owner:dns-team"}),
				TTL:      khulnasoft.F(dns.TTL1),
			},
		},
	)
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecordExport(t *testing.T) {
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
	_, err := client.DNS.Records.Export(context.TODO(), dns.RecordExportParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecordGet(t *testing.T) {
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
	_, err := client.DNS.Records.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.RecordGetParams{
			ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		},
	)
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecordImportWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.DNS.Records.Import(context.TODO(), dns.RecordImportParams{
		ZoneID:  khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		File:    khulnasoft.F("www.example.com. 300 IN  A 127.0.0.1"),
		Proxied: khulnasoft.F("true"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecordScan(t *testing.T) {
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
	_, err := client.DNS.Records.Scan(context.TODO(), dns.RecordScanParams{
		ZoneID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body:   map[string]interface{}{},
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
