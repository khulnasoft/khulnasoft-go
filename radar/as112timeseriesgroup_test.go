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

func TestAS112TimeseriesGroupDNSSECWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AS112.TimeseriesGroups.DNSSEC(context.TODO(), radar.AS112TimeseriesGroupDNSSECParams{
		AggInterval: khulnasoft.F(radar.AS112TimeseriesGroupDNSSECParamsAggInterval15m),
		ASN:         khulnasoft.F([]string{"string", "string", "string"}),
		Continent:   khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      khulnasoft.F(radar.AS112TimeseriesGroupDNSSECParamsFormatJson),
		Location:    khulnasoft.F([]string{"string", "string", "string"}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAS112TimeseriesGroupEdnsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AS112.TimeseriesGroups.Edns(context.TODO(), radar.AS112TimeseriesGroupEdnsParams{
		AggInterval: khulnasoft.F(radar.AS112TimeseriesGroupEdnsParamsAggInterval15m),
		ASN:         khulnasoft.F([]string{"string", "string", "string"}),
		Continent:   khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      khulnasoft.F(radar.AS112TimeseriesGroupEdnsParamsFormatJson),
		Location:    khulnasoft.F([]string{"string", "string", "string"}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAS112TimeseriesGroupIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AS112.TimeseriesGroups.IPVersion(context.TODO(), radar.AS112TimeseriesGroupIPVersionParams{
		AggInterval: khulnasoft.F(radar.AS112TimeseriesGroupIPVersionParamsAggInterval15m),
		ASN:         khulnasoft.F([]string{"string", "string", "string"}),
		Continent:   khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      khulnasoft.F(radar.AS112TimeseriesGroupIPVersionParamsFormatJson),
		Location:    khulnasoft.F([]string{"string", "string", "string"}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAS112TimeseriesGroupProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AS112.TimeseriesGroups.Protocol(context.TODO(), radar.AS112TimeseriesGroupProtocolParams{
		AggInterval: khulnasoft.F(radar.AS112TimeseriesGroupProtocolParamsAggInterval15m),
		ASN:         khulnasoft.F([]string{"string", "string", "string"}),
		Continent:   khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      khulnasoft.F(radar.AS112TimeseriesGroupProtocolParamsFormatJson),
		Location:    khulnasoft.F([]string{"string", "string", "string"}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAS112TimeseriesGroupQueryTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AS112.TimeseriesGroups.QueryType(context.TODO(), radar.AS112TimeseriesGroupQueryTypeParams{
		AggInterval: khulnasoft.F(radar.AS112TimeseriesGroupQueryTypeParamsAggInterval15m),
		ASN:         khulnasoft.F([]string{"string", "string", "string"}),
		Continent:   khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      khulnasoft.F(radar.AS112TimeseriesGroupQueryTypeParamsFormatJson),
		Location:    khulnasoft.F([]string{"string", "string", "string"}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAS112TimeseriesGroupResponseCodesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.AS112.TimeseriesGroups.ResponseCodes(context.TODO(), radar.AS112TimeseriesGroupResponseCodesParams{
		AggInterval: khulnasoft.F(radar.AS112TimeseriesGroupResponseCodesParamsAggInterval15m),
		ASN:         khulnasoft.F([]string{"string", "string", "string"}),
		Continent:   khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:     khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      khulnasoft.F(radar.AS112TimeseriesGroupResponseCodesParamsFormatJson),
		Location:    khulnasoft.F([]string{"string", "string", "string"}),
		Name:        khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
