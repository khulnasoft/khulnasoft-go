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

func TestAttackLayer7TimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Timeseries(context.TODO(), radar.AttackLayer7TimeseriesParams{
		AggInterval:       khulnasoft.F(radar.AttackLayer7TimeseriesParamsAggInterval15m),
		ASN:               khulnasoft.F([]string{"string", "string", "string"}),
		Attack:            khulnasoft.F([]radar.AttackLayer7TimeseriesParamsAttack{radar.AttackLayer7TimeseriesParamsAttackDDoS, radar.AttackLayer7TimeseriesParamsAttackWAF, radar.AttackLayer7TimeseriesParamsAttackBotManagement}),
		Continent:         khulnasoft.F([]string{"string", "string", "string"}),
		DateEnd:           khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         khulnasoft.F([]string{"7d", "7d", "7d"}),
		DateStart:         khulnasoft.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            khulnasoft.F(radar.AttackLayer7TimeseriesParamsFormatJson),
		HTTPMethod:        khulnasoft.F([]radar.AttackLayer7TimeseriesParamsHTTPMethod{radar.AttackLayer7TimeseriesParamsHTTPMethodGet, radar.AttackLayer7TimeseriesParamsHTTPMethodPost, radar.AttackLayer7TimeseriesParamsHTTPMethodDelete}),
		HTTPVersion:       khulnasoft.F([]radar.AttackLayer7TimeseriesParamsHTTPVersion{radar.AttackLayer7TimeseriesParamsHTTPVersionHttPv1, radar.AttackLayer7TimeseriesParamsHTTPVersionHttPv2, radar.AttackLayer7TimeseriesParamsHTTPVersionHttPv3}),
		IPVersion:         khulnasoft.F([]radar.AttackLayer7TimeseriesParamsIPVersion{radar.AttackLayer7TimeseriesParamsIPVersionIPv4, radar.AttackLayer7TimeseriesParamsIPVersionIPv6}),
		Location:          khulnasoft.F([]string{"string", "string", "string"}),
		MitigationProduct: khulnasoft.F([]radar.AttackLayer7TimeseriesParamsMitigationProduct{radar.AttackLayer7TimeseriesParamsMitigationProductDDoS, radar.AttackLayer7TimeseriesParamsMitigationProductWAF, radar.AttackLayer7TimeseriesParamsMitigationProductBotManagement}),
		Name:              khulnasoft.F([]string{"string", "string", "string"}),
		Normalization:     khulnasoft.F(radar.AttackLayer7TimeseriesParamsNormalizationPercentageChange),
	})
	if err != nil {
		var apierr *khulnasoft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
