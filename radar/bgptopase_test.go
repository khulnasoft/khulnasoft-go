// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/khulnasoft/khulnasoft-go/internal/testutil"
	"github.com/khulnasoft/khulnasoft-go/option"
	"github.com/khulnasoft/khulnasoft-go/radar"
)

func TestBGPTopAseGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.BGP.Top.Ases.Get(context.TODO(), radar.BGPTopAseGetParams{
		ASN:        khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:    khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:  khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:     khulnasoft.F(radar.BGPTopAseGetParamsFormatJson),
		Limit:      khulnasoft.F(int64(5)),
		Name:       khulnasoft.F([]string{"string", "string", "string"}),
		Prefix:     khulnasoft.F([]string{"1.1.1.0/24", "1.1.1.0/24", "1.1.1.0/24"}),
		UpdateType: khulnasoft.F([]radar.BGPTopAseGetParamsUpdateType{radar.BGPTopAseGetParamsUpdateTypeAnnouncement, radar.BGPTopAseGetParamsUpdateTypeWithdrawal}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBGPTopAsePrefixesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.BGP.Top.Ases.Prefixes(context.TODO(), radar.BGPTopAsePrefixesParams{
		Country: khulnasoft.F("NZ"),
		Format:  khulnasoft.F(radar.BGPTopAsePrefixesParamsFormatJson),
		Limit:   khulnasoft.F(int64(10)),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
