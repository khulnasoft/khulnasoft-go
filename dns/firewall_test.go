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

func TestFirewallNewWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Firewall.New(context.TODO(), dns.FirewallNewParams{
		AccountID:   khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Name:        khulnasoft.F("My Awesome DNS Firewall cluster"),
		UpstreamIPs: khulnasoft.F([]dns.UpstreamIPsParam{"192.0.2.1", "198.51.100.1", "string"}),
		AttackMitigation: khulnasoft.F(dns.AttackMitigationParam{
			Enabled:                   khulnasoft.F(true),
			OnlyWhenUpstreamUnhealthy: khulnasoft.F(false),
		}),
		DeprecateAnyRequests: khulnasoft.F(true),
		ECSFallback:          khulnasoft.F(false),
		MaximumCacheTTL:      khulnasoft.F(900.000000),
		MinimumCacheTTL:      khulnasoft.F(60.000000),
		NegativeCacheTTL:     khulnasoft.F(900.000000),
		Ratelimit:            khulnasoft.F(600.000000),
		Retries:              khulnasoft.F(2.000000),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFirewallListWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Firewall.List(context.TODO(), dns.FirewallListParams{
		AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Page:      khulnasoft.F(1.000000),
		PerPage:   khulnasoft.F(1.000000),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFirewallDelete(t *testing.T) {
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
	_, err := client.DNS.Firewall.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.FirewallDeleteParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestFirewallEditWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Firewall.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.FirewallEditParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AttackMitigation: khulnasoft.F(dns.AttackMitigationParam{
				Enabled:                   khulnasoft.F(true),
				OnlyWhenUpstreamUnhealthy: khulnasoft.F(false),
			}),
			DeprecateAnyRequests: khulnasoft.F(true),
			ECSFallback:          khulnasoft.F(false),
			MaximumCacheTTL:      khulnasoft.F(900.000000),
			MinimumCacheTTL:      khulnasoft.F(60.000000),
			Name:                 khulnasoft.F("My Awesome DNS Firewall cluster"),
			NegativeCacheTTL:     khulnasoft.F(900.000000),
			Ratelimit:            khulnasoft.F(600.000000),
			Retries:              khulnasoft.F(2.000000),
			UpstreamIPs:          khulnasoft.F([]dns.UpstreamIPsParam{"192.0.2.1", "198.51.100.1", "string"}),
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

func TestFirewallGet(t *testing.T) {
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
	_, err := client.DNS.Firewall.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.FirewallGetParams{
			AccountID: khulnasoft.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
