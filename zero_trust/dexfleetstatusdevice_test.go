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

func TestDEXFleetStatusDeviceListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DEX.FleetStatus.Devices.List(context.TODO(), zero_trust.DEXFleetStatusDeviceListParams{
		AccountID: khulnasoft.F("01a7362d577a6c3019a474fd6f485823"),
		From:      khulnasoft.F("2023-10-11T00:00:00Z"),
		Page:      khulnasoft.F(1.000000),
		PerPage:   khulnasoft.F(10.000000),
		To:        khulnasoft.F("2023-10-11T00:00:00Z"),
		Colo:      khulnasoft.F("SJC"),
		DeviceID:  khulnasoft.F("cb49c27f-7f97-49c5-b6f3-f7c01ead0fd7"),
		Mode:      khulnasoft.F("proxy"),
		Platform:  khulnasoft.F("windows"),
		SortBy:    khulnasoft.F(zero_trust.DEXFleetStatusDeviceListParamsSortByColo),
		Source:    khulnasoft.F(zero_trust.DEXFleetStatusDeviceListParamsSourceLastSeen),
		Status:    khulnasoft.F("connected"),
		Version:   khulnasoft.F("1.0.0"),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
