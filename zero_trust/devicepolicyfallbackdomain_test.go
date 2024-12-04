// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/zero_trust"
)

func TestDevicePolicyFallbackDomainUpdate(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.FallbackDomains.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DevicePolicyFallbackDomainUpdateParams{
			AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
			Body: []zero_trust.FallbackDomainParam{{
				Suffix:      khulnasoft.F("example.com"),
				Description: khulnasoft.F("Domain bypass for local development"),
				DNSServer:   khulnasoft.F([]string{"1.1.1.1", "1.1.1.1", "1.1.1.1"}),
			}, {
				Suffix:      khulnasoft.F("example.com"),
				Description: khulnasoft.F("Domain bypass for local development"),
				DNSServer:   khulnasoft.F([]string{"1.1.1.1", "1.1.1.1", "1.1.1.1"}),
			}, {
				Suffix:      khulnasoft.F("example.com"),
				Description: khulnasoft.F("Domain bypass for local development"),
				DNSServer:   khulnasoft.F([]string{"1.1.1.1", "1.1.1.1", "1.1.1.1"}),
			}},
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

func TestDevicePolicyFallbackDomainList(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.FallbackDomains.List(context.TODO(), zero_trust.DevicePolicyFallbackDomainListParams{
		AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDevicePolicyFallbackDomainGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Devices.Policies.FallbackDomains.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DevicePolicyFallbackDomainGetParams{
			AccountID: khulnasoft.F("699d98642c564d2e855e9661899b7252"),
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
