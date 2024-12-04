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

func TestAttackLayer7TopLocationOriginWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.Locations.Origin(context.TODO(), radar.AttackLayer7TopLocationOriginParams{
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7TopLocationOriginParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7TopLocationOriginParamsHTTPMethod{radar.AttackLayer7TopLocationOriginParamsHTTPMethodGet, radar.AttackLayer7TopLocationOriginParamsHTTPMethodPost, radar.AttackLayer7TopLocationOriginParamsHTTPMethodDelete}),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7TopLocationOriginParamsHTTPVersion{radar.AttackLayer7TopLocationOriginParamsHTTPVersionHttPv1, radar.AttackLayer7TopLocationOriginParamsHTTPVersionHttPv2, radar.AttackLayer7TopLocationOriginParamsHTTPVersionHttPv3}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7TopLocationOriginParamsIPVersion{radar.AttackLayer7TopLocationOriginParamsIPVersionIPv4, radar.AttackLayer7TopLocationOriginParamsIPVersionIPv6}),
		Limit:             khulnasoft.F(int64(5)),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7TopLocationOriginParamsMitigationProduct{radar.AttackLayer7TopLocationOriginParamsMitigationProductDDoS, radar.AttackLayer7TopLocationOriginParamsMitigationProductWAF, radar.AttackLayer7TopLocationOriginParamsMitigationProductBotManagement}),
		Name:              khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttackLayer7TopLocationTargetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.Locations.Target(context.TODO(), radar.AttackLayer7TopLocationTargetParams{
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7TopLocationTargetParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7TopLocationTargetParamsHTTPMethod{radar.AttackLayer7TopLocationTargetParamsHTTPMethodGet, radar.AttackLayer7TopLocationTargetParamsHTTPMethodPost, radar.AttackLayer7TopLocationTargetParamsHTTPMethodDelete}),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7TopLocationTargetParamsHTTPVersion{radar.AttackLayer7TopLocationTargetParamsHTTPVersionHttPv1, radar.AttackLayer7TopLocationTargetParamsHTTPVersionHttPv2, radar.AttackLayer7TopLocationTargetParamsHTTPVersionHttPv3}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7TopLocationTargetParamsIPVersion{radar.AttackLayer7TopLocationTargetParamsIPVersionIPv4, radar.AttackLayer7TopLocationTargetParamsIPVersionIPv6}),
		Limit:             khulnasoft.F(int64(5)),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7TopLocationTargetParamsMitigationProduct{radar.AttackLayer7TopLocationTargetParamsMitigationProductDDoS, radar.AttackLayer7TopLocationTargetParamsMitigationProductWAF, radar.AttackLayer7TopLocationTargetParamsMitigationProductBotManagement}),
		Name:              khulnasoft.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
