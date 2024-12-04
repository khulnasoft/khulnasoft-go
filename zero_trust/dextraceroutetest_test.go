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

func TestDEXTracerouteTestGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DEX.TracerouteTests.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DEXTracerouteTestGetParams{
			AccountID: khulnasoft.F("01a7362d577a6c3019a474fd6f485823"),
			From:      khulnasoft.F("1689520412000"),
			Interval:  khulnasoft.F(zero_trust.DEXTracerouteTestGetParamsIntervalMinute),
			To:        khulnasoft.F("1689606812000"),
			Colo:      khulnasoft.F("colo"),
			DeviceID:  khulnasoft.F([]string{"string", "string", "string"}),
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

func TestDEXTracerouteTestNetworkPath(t *testing.T) {
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
	_, err := client.ZeroTrust.DEX.TracerouteTests.NetworkPath(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DEXTracerouteTestNetworkPathParams{
			AccountID: khulnasoft.F("01a7362d577a6c3019a474fd6f485823"),
			DeviceID:  khulnasoft.F("deviceId"),
			From:      khulnasoft.F("1689520412000"),
			Interval:  khulnasoft.F(zero_trust.DEXTracerouteTestNetworkPathParamsIntervalMinute),
			To:        khulnasoft.F("1689606812000"),
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

func TestDEXTracerouteTestPercentilesWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DEX.TracerouteTests.Percentiles(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.DEXTracerouteTestPercentilesParams{
			AccountID: khulnasoft.F("01a7362d577a6c3019a474fd6f485823"),
			From:      khulnasoft.F("2023-09-20T17:00:00Z"),
			To:        khulnasoft.F("2023-09-20T17:00:00Z"),
			Colo:      khulnasoft.F("colo"),
			DeviceID:  khulnasoft.F([]string{"string", "string", "string"}),
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
