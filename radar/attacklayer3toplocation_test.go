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

func TestAttackLayer3TopLocationOriginWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.Locations.Origin(context.TODO(), radar.AttackLayer3TopLocationOriginParams{
		Continent: khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    khulnasoft.F(radar.AttackLayer3TopLocationOriginParamsFormatJson),
		IPVersion: khulnasoft.F([]radar.AttackLayer3TopLocationOriginParamsIPVersion{radar.AttackLayer3TopLocationOriginParamsIPVersionIPv4, radar.AttackLayer3TopLocationOriginParamsIPVersionIPv6}),
		Limit:     khulnasoft.F(int64(5)),
		Location:  khulnasoft.F([]string{"string", "string", "string"}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
		Protocol:  khulnasoft.F([]radar.AttackLayer3TopLocationOriginParamsProtocol{radar.AttackLayer3TopLocationOriginParamsProtocolUdp, radar.AttackLayer3TopLocationOriginParamsProtocolTCP, radar.AttackLayer3TopLocationOriginParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer3TopLocationTargetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.Locations.Target(context.TODO(), radar.AttackLayer3TopLocationTargetParams{
		Continent: khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:   khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart: khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:    khulnasoft.F(radar.AttackLayer3TopLocationTargetParamsFormatJson),
		IPVersion: khulnasoft.F([]radar.AttackLayer3TopLocationTargetParamsIPVersion{radar.AttackLayer3TopLocationTargetParamsIPVersionIPv4, radar.AttackLayer3TopLocationTargetParamsIPVersionIPv6}),
		Limit:     khulnasoft.F(int64(5)),
		Location:  khulnasoft.F([]string{"string", "string", "string"}),
		Name:      khulnasoft.F([]string{"string", "string", "string"}),
		Protocol:  khulnasoft.F([]radar.AttackLayer3TopLocationTargetParamsProtocol{radar.AttackLayer3TopLocationTargetParamsProtocolUdp, radar.AttackLayer3TopLocationTargetParamsProtocolTCP, radar.AttackLayer3TopLocationTargetParamsProtocolIcmp}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
