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

func TestAttackLayer3SummaryBitrateWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.Bitrate(context.TODO(), radar.AttackLayer3SummaryBitrateParams{
		Continent: khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: khulnasoft.F(radar.AttackLayer3SummaryBitrateParamsDirectionOrigin),
		Format:    khulnasoft.F(radar.AttackLayer3SummaryBitrateParamsFormatJson),
		IPVersion: khulnasoft.F([]radar.AttackLayer3SummaryBitrateParamsIPVersion{radar.AttackLayer3SummaryBitrateParamsIPVersionIPv4, radar.AttackLayer3SummaryBitrateParamsIPVersionIPv6}),
		Location:  khulnasoft.F([]string{"string", "string", "string"}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
		Protocol:  khulnasoft.F([]radar.AttackLayer3SummaryBitrateParamsProtocol{radar.AttackLayer3SummaryBitrateParamsProtocolUdp, radar.AttackLayer3SummaryBitrateParamsProtocolTCP, radar.AttackLayer3SummaryBitrateParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3SummaryDurationWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.Duration(context.TODO(), radar.AttackLayer3SummaryDurationParams{
		Continent: khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: khulnasoft.F(radar.AttackLayer3SummaryDurationParamsDirectionOrigin),
		Format:    khulnasoft.F(radar.AttackLayer3SummaryDurationParamsFormatJson),
		IPVersion: khulnasoft.F([]radar.AttackLayer3SummaryDurationParamsIPVersion{radar.AttackLayer3SummaryDurationParamsIPVersionIPv4, radar.AttackLayer3SummaryDurationParamsIPVersionIPv6}),
		Location:  khulnasoft.F([]string{"string", "string", "string"}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
		Protocol:  khulnasoft.F([]radar.AttackLayer3SummaryDurationParamsProtocol{radar.AttackLayer3SummaryDurationParamsProtocolUdp, radar.AttackLayer3SummaryDurationParamsProtocolTCP, radar.AttackLayer3SummaryDurationParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3SummaryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.Get(context.TODO(), radar.AttackLayer3SummaryGetParams{
		ASN:       khulnasoft.F([]string{"string", "string", "string"}),
		Continent: khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    khulnasoft.F(radar.AttackLayer3SummaryGetParamsFormatJson),
		Location:  khulnasoft.F([]string{"string", "string", "string"}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3SummaryIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.IPVersion(context.TODO(), radar.AttackLayer3SummaryIPVersionParams{
		Continent: khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: khulnasoft.F(radar.AttackLayer3SummaryIPVersionParamsDirectionOrigin),
		Format:    khulnasoft.F(radar.AttackLayer3SummaryIPVersionParamsFormatJson),
		Location:  khulnasoft.F([]string{"string", "string", "string"}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
		Protocol:  khulnasoft.F([]radar.AttackLayer3SummaryIPVersionParamsProtocol{radar.AttackLayer3SummaryIPVersionParamsProtocolUdp, radar.AttackLayer3SummaryIPVersionParamsProtocolTCP, radar.AttackLayer3SummaryIPVersionParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3SummaryProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.Protocol(context.TODO(), radar.AttackLayer3SummaryProtocolParams{
		Continent: khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: khulnasoft.F(radar.AttackLayer3SummaryProtocolParamsDirectionOrigin),
		Format:    khulnasoft.F(radar.AttackLayer3SummaryProtocolParamsFormatJson),
		IPVersion: khulnasoft.F([]radar.AttackLayer3SummaryProtocolParamsIPVersion{radar.AttackLayer3SummaryProtocolParamsIPVersionIPv4, radar.AttackLayer3SummaryProtocolParamsIPVersionIPv6}),
		Location:  khulnasoft.F([]string{"string", "string", "string"}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3SummaryVectorWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.Vector(context.TODO(), radar.AttackLayer3SummaryVectorParams{
		Continent: khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Direction: khulnasoft.F(radar.AttackLayer3SummaryVectorParamsDirectionOrigin),
		Format:    khulnasoft.F(radar.AttackLayer3SummaryVectorParamsFormatJson),
		IPVersion: khulnasoft.F([]radar.AttackLayer3SummaryVectorParamsIPVersion{radar.AttackLayer3SummaryVectorParamsIPVersionIPv4, radar.AttackLayer3SummaryVectorParamsIPVersionIPv6}),
		Location:  khulnasoft.F([]string{"string", "string", "string"}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
		Protocol:  khulnasoft.F([]radar.AttackLayer3SummaryVectorParamsProtocol{radar.AttackLayer3SummaryVectorParamsProtocolUdp, radar.AttackLayer3SummaryVectorParamsProtocolTCP, radar.AttackLayer3SummaryVectorParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
